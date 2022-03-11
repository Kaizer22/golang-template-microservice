package impl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"main/logging"
	"main/model/entity"
	"main/repository"
	"main/utils"
)

func NewPgPromoRepository(connection *sql.DB) repository.PromoRepository {
	return pgPromoRepository{
		db: connection,
	}
}

type pgPromoRepository struct {
	db *sql.DB
}

var insQ = "INSERT INTO promotion (name, description) VALUES ($1, $2) RETURNING id;"

func (p pgPromoRepository) InsertPromo(info entity.PromoInfo) (int, error) {
	lastInsertId := 0
	err := utils.RunWithProfiler(repository.TagInsPromo,
		func() error {
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			err = tx.QueryRow(insQ,
				info.Name,
				info.Description,
			).Scan(&lastInsertId)
			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

const getAllPrQ = "SELECT id, name, description FROM promotion"

func (p pgPromoRepository) GetAllPromos() ([]*entity.PromoShort, error) {
	var res []*entity.PromoShort
	err := utils.RunWithProfiler(repository.TagGetAllPromo,
		func() error {
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.Prepare(getAllPrQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			result, err := stmt.Query()
			if err != nil {
				logging.ErrorFormat("could not get all products: %s", err)
			}
			for result.Next() {
				promoShort := entity.PromoShort{}
				err := result.Scan(
					&promoShort.Id,
					&promoShort.Name,
					&promoShort.Description,
				)
				if err != nil {
					logging.ErrorFormat("could not read product: %s", err)
				}
				res = append(res, &promoShort)
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}

const getPromoByIdQ = "SELECT id, name, description, array_to_json(participants), array_to_json(prizes) " +
	"FROM promotion WHERE id = $1"

func (p pgPromoRepository) GetPromoById(id int) (*entity.Promotion, error) {
	res := entity.Promotion{}
	err := utils.RunWithProfiler(repository.TagGetPromoById,
		func() error {
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.Prepare(getPromoByIdQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			row := stmt.QueryRow(id)

			var participants string
			var prizes string
			var prtIDs []int
			var przIds []int
			err = row.Scan(
				&res.Id,
				&res.Name,
				&res.Description,
				&res.Participants,
				prizes,
				participants,
			)
			err = json.Unmarshal([]byte(participants), &prtIDs)
			if err != nil {
				return err
			}
			err = json.Unmarshal([]byte(prizes), &przIds)
			if err != nil {
				return err
			}
			prtRes, _ := p.GetParticipantsByIds(prtIDs)
			przRes, _ := p.GetPrizesByIds(przIds)

			res.Prizes = przRes
			res.Participants = prtRes

			if err != nil {
				logging.ErrorFormat("could not get product by id %d ", id)
				return err
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

const getParticipantsByIds = "SELECT id, name FROM participant WHERE id IN $1"

func (p pgPromoRepository) GetParticipantsByIds(ids []int) ([]*entity.Participant, error) {
	res := []*entity.Participant{}
	tx, err := p.db.Begin()
	if err != nil {
		logging.Error("could not begin a transaction")
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(getParticipantsByIds)
	if err != nil {
		logging.Error("could not prepare a statement")
		return nil, err
	}

	result, err := stmt.Query(pq.Array(ids))
	if err != nil {
		logging.ErrorFormat("could not get all products: %s", err)
	}
	for result.Next() {
		prt := entity.Participant{}
		err := result.Scan(
			&prt.Id,
			&prt.Name,
		)
		if err != nil {
			logging.ErrorFormat("could not read product: %s", err)
		}
		res = append(res, &prt)
	}

	if err = tx.Commit(); err != nil {
		logging.Error("could not commit a transaction")
		return nil, err
	}
	return res, nil
}

const getPrizesByIds = "SELECT id, description FROM prize WHERE id IN $1"

func (p pgPromoRepository) GetPrizesByIds(ids []int) ([]*entity.Prize, error) {
	res := []*entity.Prize{}
	tx, err := p.db.Begin()
	if err != nil {
		logging.Error("could not begin a transaction")
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(getParticipantsByIds)
	if err != nil {
		logging.Error("could not prepare a statement")
		return nil, err
	}

	result, err := stmt.Query(pq.Array(ids))
	if err != nil {
		logging.ErrorFormat("could not get all products: %s", err)
	}
	for result.Next() {
		prt := entity.Prize{}
		err := result.Scan(
			&prt.Id,
			&prt.Description,
		)
		if err != nil {
			logging.ErrorFormat("could not read product: %s", err)
		}
		res = append(res, &prt)
	}

	if err = tx.Commit(); err != nil {
		logging.Error("could not commit a transaction")
		return nil, err
	}
	return res, nil
}

const updQ = "UPDATE promotion SET name = $1, description = $2" +
	"WHERE id = $3"

func (p pgPromoRepository) UpdatePromo(id int, info entity.PromoInfo) error {
	err := utils.RunWithProfiler(repository.TagUpdPromo,
		func() error {
			if len(info.Name) == 0 {
				return errors.New("name cannot be empty")
			}
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.Prepare(updQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			_, err = stmt.Exec(
				info.Name,
				info.Description,
				id,
			)
			if err != nil {
				logging.ErrorFormat("could not update product %d %s", id, info.Name)
				return err
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

const delQ = "DELETE FROM promotion WHERE id = $1"

func (p pgPromoRepository) DeletePromo(promoId int) error {
	err := utils.RunWithProfiler(repository.TagDelPromo,
		func() error {
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.Prepare(delQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			_, err = stmt.Exec(promoId)
			if err != nil {
				logging.ErrorFormat("could not delete promo %d %s", promoId)
				return err
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

const updParticipantsQ = "UPDATE promotion SET participants = $1 WHERE id = $2"
const updPrizesQ = "UPDATE promotion SET prizes = $1 WHERE id = $2"

const isnParticipantQ = "INSERT INTO participant(name) VALUES ($1) RETURNING (i)"

func (p pgPromoRepository) AddParticipant(promoId int, pInfo entity.ParticipantInfo) (int, error) {
	var lastInsertId int
	err := utils.RunWithProfiler(repository.TagInsPromoParticipant,
		func() error {
			promotion, err := p.GetPromoById(promoId)
			if err != nil {
				logging.Error("could not get a promo")
				return err
			}
			currParticipants := promotion.Participants
			tx, err := p.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			err = tx.QueryRow(insQ,
				pInfo.Name,
			).Scan(&lastInsertId)
			if err != nil {
				logging.ErrorFormat("could not add participant to %d %s", promoId, err)
				return err
			}

			currParticipants = append(currParticipants, &entity.Participant{
				Id:   lastInsertId,
				Name: pInfo.Name,
			})

			stmt, err := tx.Prepare(updParticipantsQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			_, err = stmt.Exec(promoId, pq.Array(getPrtIdsArray(currParticipants)), promoId)
			if err != nil {
				logging.ErrorFormat("could not update participants for %d %s", promoId, err)
				return err
			}
			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func getPrtIdsArray(arr []*entity.Participant) (res []int) {
	for _, participant := range arr {
		res = append(res, participant.Id)
	}
	return
}

func getPrzIdsArray(arr []*entity.Prize) (res []int) {
	for _, prize := range arr {
		res = append(res, prize.Id)
	}
	return
}

func (p pgPromoRepository) DeleteParticipant(promoId int, participantId int) error {
	//TODO implement me
	return errors.New("implement me")
}

const isnPrizeQ = "INSERT INTO prize(name) VALUES ($1) RETURNING (id)"

func (p pgPromoRepository) AddPrize(promoId int, inf entity.PrizeInfo) (int, error) {
	//TODO implement me
	return 0, errors.New("implement me")
}

func (p pgPromoRepository) DeletePrize(promoId int, prizeId int) error {
	//TODO implement me
	return errors.New("implement me")
}

func (p pgPromoRepository) GetWinners(promoId int) ([]*entity.PromoResult, error) {
	//TODO implement me
	return nil, errors.New("implement me")
}
