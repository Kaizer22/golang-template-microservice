package main

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/logging"
	"main/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

const (
	apiV1 = "/api/v1"
)

func main() {
	r := gin.Default()
	c := controller.NewProductController()

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
	initService()
	logging.InfoFormat("Starting server at %s", addr)
	err := r.Run(addr)
	if err != nil {
		logging.FatalFormat("unable to start server")
		panic(err)
		return
	}

}

func initService() {
	addr = utils.GetEnv(utils.ListenAddressEnvKey, ":8080")

}
