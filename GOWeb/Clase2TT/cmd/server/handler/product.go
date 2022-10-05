package handler

import (
	"strconv"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase2TT/internal/products"
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
		if err := c.Bind(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
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
