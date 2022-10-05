package main

import (
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase2TT/cmd/server/handler"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase2TT/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	req := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", req.Store())
	pr.GET("/", req.GetAll())
	pr.PUT("/:id", req.Update())
	r.Run()
}
