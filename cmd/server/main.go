package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlerPing "github.com/jum8/EBE3_GoWeb.git/cmd/server/handler/ping"
	handlerProduct "github.com/jum8/EBE3_GoWeb.git/cmd/server/handler/product"
	"github.com/jum8/EBE3_GoWeb.git/internal/product"
	"github.com/jum8/EBE3_GoWeb.git/pkg/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/jum8/EBE3_GoWeb.git/docs"
)

const (
	port = ":8080"
)


// @title Certified Tech Developer
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://developers.ctd.com.ar/es ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	controllerPing := handlerPing.NewControllerPing()
	repo := product.NewMemoryRespository()
	service := product.NewProductService(repo)
	productController := handlerProduct.NewProductController(service)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group := engine.Group("api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		productsGroup := group.Group("/products")
		{
			productsGroup.GET("", productController.HandlerGetAll())
			productsGroup.GET("/:id", productController.HandlerGetById())
			productsGroup.POST("", middleware.Authenticate(), productController.HandlerSaveProduct())
			productsGroup.PUT("/:id", middleware.Authenticate(), productController.HandlerUpdateProduct())
			productsGroup.DELETE("/:id", middleware.Authenticate(), productController.HandlerDeleteProduct())
		}
		
	}


	if err := engine.Run(port); err != nil {
		log.Fatal(err)
	}

}
