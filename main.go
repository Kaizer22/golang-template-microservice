package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/db"
	conn "main/db/impl"
	"main/logging"
	repo "main/repository/impl"
	"main/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs"
)

var (
	addr string
)

// @title           API для промоакций и розыгрыша призов
// @version         1.0
// @description     RESTful API
// @termsOfService  http://swagger.io/terms/

// @contact.name   GitHub repository
// @contact.url    http://www.swagger.io/support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /promo

// @securityDefinitions.basic  BasicAuth

const (
	promo = "/promo"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	connection, err := conn.NewPgConnectionProvider()
	if err != nil {
		panic(err)
	}
	err = connection.Migrate(db.PgMigrationsPath)
	if err != nil {
		panic(err)
	}

	promoRepo := repo.NewPgPromoRepository(connection.Connection())
	c := controller.NewPromoController(ctx, promoRepo)

	v1Promo := r.Group(promo)
	{
		v1Promo.POST("", c.AddPromo)
		v1Promo.GET("", c.ListPromos)
		v1Promo.GET(":promoId", c.GetPromo)
		v1Promo.PUT(":promoId", c.PutPromo)
		v1Promo.DELETE(":promoId", c.DeletePromo)
		v1Promo.POST(":promoId/participant", c.AddParticipant)
		v1Promo.DELETE(":promoId/participant/:participantId", c.DeleteParticipant)
		v1Promo.POST("/:promoId/prize", c.DeletePrize)
		v1Promo.DELETE(":promoId/prize/:prizeId", c.DeletePrize)
		v1Promo.POST(":promoId/raffle", c.GetWinners)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	initService()
	logging.InfoFormat("Starting server at %s", addr)
	err = r.Run(addr)
	if err != nil {
		logging.FatalFormat("unable to start server")
		panic(err)
		return
	}

}

func initService() {
	addr = utils.GetEnv(utils.ListenAddressEnvKey, ":8080")

}
