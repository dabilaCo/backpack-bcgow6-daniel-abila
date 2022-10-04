package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generamos una estructura con los campos de la petición que vamos a recibir
type request struct {
	ID     int     `json:"id"`
	Name   string  `json:"name" binding:"required"` //binding:"required" para indicar que ese campo es requerido
	Type   string  `json:"type" binding:"required"`
	Amount int     `json:"amount" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

// Generamos una lista para ir guardando los datos que vamos recibiendo
var arrayProducts []request

func main() {
	//definimos un servicio web con metodo POST que tendrá products como path
	r := gin.Default()
	r.GET("/ping", Ping())

	pr := r.Group("products")
	pr.POST("", Store())

	r.Run()
}

// el store nos va a guardar los productos que vamos recibiendo y va incrementando el ID
func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "4230" || token == "" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada.",
			})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {

			if err := c.ShouldBindJSON(&req); err != nil { //GET ALL THE STRUCT USER BY BODY
				if req.Name == "" || req.Type == "" || req.Amount == 0 || req.Price == 0 {
					if req.Name == "" {
						c.JSON(http.StatusBadRequest, gin.H{"error": "El campo Name es requerido!"})
					}
					if req.Type == "" {
						c.JSON(http.StatusBadRequest, gin.H{"error": "El campo Type es requerido!"})
					}
					if req.Amount == 0{
						c.JSON(http.StatusBadRequest, gin.H{"error": "El campo Amount es requerido!"})
					}
					if req.Price == 0{
						c.JSON(http.StatusBadRequest, gin.H{"error": "El campo price es requerido!"})
					}
					return
				}else{

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
				}
			}
		}

		req.ID = len(arrayProducts) + 1
		arrayProducts = append(arrayProducts, req)
		c.JSON(http.StatusOK, req)


	}
}
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
