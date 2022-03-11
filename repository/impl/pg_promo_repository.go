package impl

import (
	"database/sql"
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

const getAllPrQ = "SELECT id, name, description FROM promotions"

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

func (p pgPromoRepository) GetPromoById(id int) (*entity.Promotion, error) {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) UpdatePromo(id int, info entity.PromoInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) DeletePromo(promoId int) error {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) AddParticipant(promoId int, pInfo entity.ParticipantInfo) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) DeleteParticipant(promoId int, participantId int) error {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) AddPrize(promoId int, inf entity.PrizeInfo) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) DeletePrize(promoId int, prizeId int) error {
	//TODO implement me
	panic("implement me")
}

func (p pgPromoRepository) GetWinners(promoId int) ([]*entity.PromoResult, error) {
	//TODO implement me
	panic("implement me")
}
