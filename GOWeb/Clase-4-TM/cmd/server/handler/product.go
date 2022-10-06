package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-4-TM/internal/products"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOWeb/Clase-4-TM/pkg/web"

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
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		p, err := pr.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (pr *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "el nombre del producto es requerido"))
			return
		}
		if req.TypeRquest == "" {
			c.JSON(400, web.NewResponse(400, nil, "el tipo de producto es requerido"))
			return
		}
		if req.Amount == 0 {
			c.JSON(400, web.NewResponse(400, nil, "La cantidad de unidades es requerida"))
			return
		}
		if req.Price == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "el precio es requerido"))
			return
		}

		p, err := pr.service.Store(req.Name, req.TypeRquest, req.Amount, req.Price)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

func (pr *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "ID no encontrado"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.TypeRquest == "" {
			c.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}
		if req.Amount == 0 {
			c.JSON(400, web.NewResponse(400, nil, "La cantidad de unidades es necesaria."))
			return
		}
		if req.Price == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El precio no es válido"))
			return
		}
		p, err := pr.service.Update(int(id), req.Name, req.TypeRquest, req.Amount, req.Price)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

func (pr *Product) Delete(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token != "123456" || token == "" {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido"))
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, web.NewResponse(400, nil, "ID inválido"))
		return
	}
	err = pr.service.Delete(int(id))
	if err != nil {
		c.JSON(404, web.NewResponse(404, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, web.NewResponse(200, "El producto ha sido eliminado", ""))
}

func (pr *Product) UpdateNameAndPrice(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token != "123456" || token == "" {
		c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido"))
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, web.NewResponse(400, nil, "ID inválido"))
		return
	}
	var request request
	if err := c.ShouldBindJSON(&request); err != nil {
		if request.Name == "" {
			c.JSON(400, web.NewResponse(400, nil, "Nombre inválido"))
			return
		}
		if request.Price == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Precio inválido"))
			return
		}
		return
	}
	productToUpdate, err := pr.service.UpdateNameAndPrice(int(id), request.Name, request.Price)
	if err != nil {
		c.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, web.NewResponse(200, productToUpdate, ""))
}
