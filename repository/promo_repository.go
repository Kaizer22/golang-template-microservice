package repository

import "main/model/entity"

const (
	TagInsPromo            = "INSERT PROMO"
	TagGetAllPromo         = "GET ALL PROMOS"
	TagGetPromoById        = "GET PROMO BY ID"
	TagUpdPromo            = "UPDATE PROMO"
	TagDelPromo            = "DELETE PROMO"
	TagInsPromoParticipant = "INSERT PARTICIPANT"
	TagDelPromoParticipant = "DELETE PARTICIPANT"
	TagInsPromoPrize       = "INSERT PRIZE"
	TagDelPromoPrize       = "DELETE PRIZE"
	TagPromoRaffle         = "GET RESULTS"
)

type PromoRepository interface {
	InsertPromo(info entity.PromoInfo) (int, error)
	GetAllPromos() ([]*entity.PromoShort, error)
	GetPromoById(id int) (*entity.Promotion, error)
	UpdatePromo(id int, info entity.PromoInfo) error
	DeletePromo(promoId int) error
	AddParticipant(promoId int, pInfo entity.ParticipantInfo) (int, error)
	DeleteParticipant(promoId int, participantId int) error
	AddPrize(promoId int, inf entity.PrizeInfo) (int, error)
	DeletePrize(promoId int, prizeId int) error
	GetWinners(promoId int) ([]*entity.PromoResult, error)
}
