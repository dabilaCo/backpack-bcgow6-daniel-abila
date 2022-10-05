package main

import (
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TM/cmd/server/handler"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TM/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
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
