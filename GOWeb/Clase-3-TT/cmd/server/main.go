package main

import (
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TT/cmd/server/handler"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TT/internal/products"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TT/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	req := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", req.Store()) //crear producto
	pr.GET("/", req.GetAll())
	pr.DELETE("/:id", req.Delete)
	pr.PUT("/:id", req.Update())
	pr.PATCH("/:id", req.UpdateNameAndPrice)

	r.Run()
}
