package main

import (
	"log"
	"os"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/cmd/server/handler"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/docs"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/internal/products"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/pkg/store"
	

	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

//Documentación con Swagger

// @title	Bootcamp Go Wave 6 - API
// @version	1.0
// @description	This is a simple API development by Daniel Abila
// @termsOfService	https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name	API Support Daniel Abila
// @contact.url	http://www.swagger.io/support
// @contact.email daniel.abila@mercadolibre.com
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	req := handler.NewProduct(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.POST("/", req.Store()) //crear producto
	pr.GET("/", req.GetAll())
	pr.DELETE("/:id", req.Delete())
	pr.PUT("/:id", req.Update())
	pr.PATCH("/:id", req.UpdateNameAndPrice)

	err = r.Run()
	if err != nil{
		log.Fatal("error al iniciar.")
	}
}
