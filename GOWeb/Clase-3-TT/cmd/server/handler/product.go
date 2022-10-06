package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-3-TT/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name       string  `json:"name"`
	TypeRquest string  `json:"typeRequest"`
	Amount     int     `json:"amount"`
	Price      float64 `json:"price"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (pr *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" || token == "" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := pr.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, p)
	}
}

func (pr *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println(req.Name)
		p, err := pr.service.Store(req.Name, req.TypeRquest, req.Amount, req.Price)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}

func (pr *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			c.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.TypeRquest == "" {
			c.JSON(400, gin.H{"error": "El tipo del producto es requerido"})
			return
		}
		if req.Amount == 0 {
			c.JSON(400, gin.H{"error": "La cantidad es requerida"})
			return
		}
		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "El precio es requerido"})
			return
		}
		p, err := pr.service.Update(int(id), req.Name, req.TypeRquest, req.Amount, req.Price)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}

func (pr *Product) Delete(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token != "123456" || token == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar esta acción"})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = pr.service.Delete(int(id))
	if err != nil{
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
}

func (pr *Product) UpdateNameAndPrice(c *gin.Context){
	token := c.Request.Header.Get("token")
	if token != "123456" || token == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No tiene permisos para realizar esta petición"})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil{
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	var request request
	if err := c.ShouldBindJSON(&request); err != nil{
		if request.Name == ""{
			c.JSON(400, gin.H{"error": "el campo Nombre es requerido"})
			return
		}
		if request.Price == 0{
			c.JSON(400, gin.H{"error": "el campo Precio es requerido"})
			return
		}
		return
	}
	productToUpdate, err := pr.service.UpdateNameAndPrice(int(id), request.Name, request.Price)
	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, productToUpdate)
}