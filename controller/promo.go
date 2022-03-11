package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"main/model/entity"
	"main/repository"
	"main/utils"
	"net/http"
	"strconv"
)

type PromoController struct {
	PromoRepo repository.PromoRepository
	ctx       context.Context
}

// NewProductController example
func NewPromoController(ctx context.Context, repo repository.PromoRepository) *PromoController {
	return &PromoController{
		PromoRepo: repo,
		ctx:       ctx,
	}
}

// AddPromo godoc
// @Summary            Add new promotion
// @Description    	   Add new promo and get entity with ID in a response
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promo_info   body            entity.PromoInfo true  "Promo info"
// @Success        201  {integer} string         "Promo ID"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router                        /promo [post]
func (c *PromoController) AddPromo(ctx *gin.Context) {
	var p entity.PromoInfo
	if err := ctx.ShouldBindJSON(&p); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	id, err := c.PromoRepo.InsertPromo(p)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.String(http.StatusOK, "%d", id)
}

// ListPromos godoc
// @Summary            Get promos
// @Description   	 Returns short info about all the promos in system
// @Tags                       promos
// @Accept                     json
// @Produce                    json
// @Success        200         {array}            entity.PromoShort
// @Failure        400         {object}           utils.HTTPError
// @Failure        404         {object}           utils.HTTPError
// @Failure        500         {object}           utils.HTTPError
// @Router                     /promo [get]
func (c *PromoController) ListPromos(ctx *gin.Context) {
	if promos, err := c.PromoRepo.GetAllPromos(); err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	} else {
		ctx.JSON(http.StatusOK, promos)
		return
	}
}

// GetPromo godoc
// @Summary            Get product by ID
// @Description    Returns product by ID
// @Tags                     promos
// @Accept                   json
// @Produce                  json
// @Param               promoId             path        int  true  "Promo ID"
// @Success        200       {object}  entity.Promotion
// @Failure        400       {object}  utils.HTTPError
// @Failure        404       {object}  utils.HTTPError
// @Failure        500       {object}  utils.HTTPError
// @Router                   /promo/{promoId} [get]
func (c *PromoController) GetPromo(ctx *gin.Context) {
	id := ctx.Param("promoId")
	iID, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	promo, err := c.PromoRepo.GetPromoById(iID)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, promo)
}

// PutPromo godoc
// @Summary            Edit promotion
// @Description    Edit existing promotion
// @Tags                      promos
// @Accept                    json
// @Produce                   json
// @Param          promoId         path                        int         true  "Promotion ID"
// @Param          data       body              entity.PromoInfo  true  "Promotion info entity"
// @Success        200   	  {string}  string  "Promo updated"
// @Failure        400        {object}            utils.HTTPError
// @Failure        404        {object}            utils.HTTPError
// @Failure        500        {object}            utils.HTTPError
// @Router                    /promo/{promoId} [put]
func (c *PromoController) PutPromo(ctx *gin.Context) {
	id := ctx.Param("promoId")
	iID, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var info entity.PromoInfo
	if err := ctx.ShouldBindJSON(&info); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.PromoRepo.UpdatePromo(iID, info)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Promo updated")
}

// DeletePromo godoc
// @Summary            Delete promotion
// @Description    Delete selected promotion
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId                  path    int    true  "Promotion ID"
// @Success         200       {string}  string  "Promotion deleted"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router                       /promo/{promoId} [delete]
func (c *PromoController) DeletePromo(ctx *gin.Context) {
	promoId := ctx.Param("promoId")
	iID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.PromoRepo.DeletePromo(iID)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Promotion deleted")
}

// AddParticipant godoc
// @Summary            Add new participant
// @Description    Add new participant and get entity with ID in a response
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId                  path    int    true  "Promotion ID"
// @Param               participant   body            entity.ParticipantInfo  true  "Participant info"
// @Success        201  {integer} string         "Participant ID"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router             /promo/{promoId}/participant [post]
func (c *PromoController) AddParticipant(ctx *gin.Context) {
	promoId := ctx.Param("promoId")
	iID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var p entity.ParticipantInfo
	if err := ctx.ShouldBindJSON(&p); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	participantId, err := c.PromoRepo.AddParticipant(iID, p)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	//TODO return struct with ID
	ctx.String(http.StatusCreated, "%d", participantId)
}

// DeleteParticipant godoc
// @Summary            Delete participant from promo
// @Description   	   Delete participant from promo by ID
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId                  path    int    true  "Promotion ID"
// @Param               participantId            path    int    true  "Participant ID"
// @Success        200     {string}  string                        "Participant successfully deleted"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router                        /promo/{promoId}/participant/{participantId} [delete]
func (c *PromoController) DeleteParticipant(ctx *gin.Context) {
	promoId := ctx.Param("promoId")
	promoIID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	participantId := ctx.Param("participantId")
	participantIID, err := strconv.Atoi(participantId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.PromoRepo.DeleteParticipant(promoIID, participantIID)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Participant successfully deleted")
}

// AddPrize godoc
// @Summary            Add new prize
// @Description    Add new prize to a promo
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId            path    int    true  "Promotion ID"
// @Param               prize   body            entity.PrizeInfo true  "Prize info"
// @Success        201  {integer} string  "Prize ID"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router                /promo/{promoId}/prize [post]
func (c *PromoController) AddPrize(ctx *gin.Context) {
	var p entity.PrizeInfo
	if err := ctx.ShouldBindJSON(&p); err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	promoId := ctx.Param("promoId")
	promoIID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	prizeId, err := c.PromoRepo.AddPrize(promoIID, p)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.String(http.StatusCreated, "%s", prizeId)
}

// DeletePrize godoc
// @Summary            Delete prize from promo
// @Description   		Delete prize from promo by ID
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId            path    int    true  "Promotion ID"
// @Param               prizeId            path    int    true  "Prize ID"
// @Success        200             {string}  string                        "Prize deleted"
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router          /promo/{promoId}/prize/{prizeId} [delete]
func (c *PromoController) DeletePrize(ctx *gin.Context) {
	promoId := ctx.Param("promoId")
	promoIID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	participantId := ctx.Param("prizeId")
	prizeIID, err := strconv.Atoi(participantId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.PromoRepo.DeletePrize(promoIID, prizeIID)
	if err != nil {
		utils.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Prize deleted")
}

// GetWinners godoc
// @Summary            Get promo winners
// @Description    	   Get promo winners with prizes (only if number of participants is equals to number of prizes
// @Tags                          promos
// @Accept                        json
// @Produce                       json
// @Param               promoId            path    int    true  "Promotion ID"
// @Success        201  {array}  entity.PromoResult
// @Failure        400  {object}  utils.HTTPError
// @Failure        404  {object}  utils.HTTPError
// @Failure		   409  {object}  utils.HTTPError
// @Failure        500  {object}  utils.HTTPError
// @Router             /promo/{promoId}/raffle [post]
func (c *PromoController) GetWinners(ctx *gin.Context) {
	promoId := ctx.Param("promoId")
	promoIID, err := strconv.Atoi(promoId)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := c.PromoRepo.GetWinners(promoIID)
	if err != nil {
		utils.NewError(ctx, http.StatusConflict, err)
		return
	}
	//TODO return struct with ID
	ctx.JSON(http.StatusOK, result)
}
