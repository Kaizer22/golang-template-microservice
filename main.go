package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"main/controller"
	conn "main/db/impl"
	"main/logging"
	repo "main/repository/impl"
	"main/utils"

	"github.com/gin-contrib/pprof"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs"
)

var (
	addr string
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

const (
	apiV1 = "/api/v1"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	connection, err := conn.NewPgConnectionProvider()
	if err != nil {
		panic(err)
	}

	productRepo := repo.NewPgProductRepository(connection.Connection())
	c := controller.NewProductController(ctx, productRepo)

	v1 := r.Group(apiV1)
	{
		products := v1.Group("/products")
		{
			products.POST("", c.AddProduct)
			products.GET("", c.ListProducts)
			products.GET(":id", c.GetProduct)
			products.PUT(":id", c.PutProduct)
			products.DELETE(":id", c.DeleteProduct)

		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pprof.Register(r, "/debug/pprof")

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
