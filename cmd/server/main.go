package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlerPing "github.com/jum8/EBE3_GoWeb.git/cmd/server/handler/ping"
	handlerProduct "github.com/jum8/EBE3_GoWeb.git/cmd/server/handler/product"
	"github.com/jum8/EBE3_GoWeb.git/docs"
	"github.com/jum8/EBE3_GoWeb.git/internal/product"
	"github.com/jum8/EBE3_GoWeb.git/pkg/middleware"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	db := connectDB()

	controllerPing := handlerPing.NewControllerPing()
	repo := product.NewSqlRespository(db)
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

func connectDB() *sql.DB {
	var username, password, hostName, port, database string

	username = "root"
	password = "root"
	hostName = "localhost"
	port = "3306"
	database = "my_db"

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostName, port, database)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
