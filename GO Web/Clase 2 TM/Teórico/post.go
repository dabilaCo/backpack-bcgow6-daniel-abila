package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio" binding:"required"`
}

var (
	products []ProductRequest
)

func main() {
	r := gin.Default()
	r.GET("/ping", Ping())

	pr := r.Group("products")
	pr.POST("", Store())

	r.Run()
}

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req.Id = len(products) + 1
		products = append(products, req)
		c.JSON(http.StatusOK, req)
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
