package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/cmd/server/handler"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/internal/products"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Levantamos el sv, hay que hacer esto para ello
func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.PATCH("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}

// Genera un Request y Response
func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_UpdateProduct(t *testing.T) {
	//creamos server y definimos rutas
	r := createServer()
	//creamos Request del tipo PATCH y Response para obtener el resultado / le pasamos lo que tendrá el body, un producto
	req, rr := createRequestTest(http.MethodPost, "/products/",
		`{			
			"name": "Monitor",
			"typeRequest": "ACH-VC5",
			"amount": 26,
			"price": 1500
		}`)

	//indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	//creamos el producto actualizado
	req, rr = createRequestTest(http.MethodPatch, "/products/1",
		`{
				"name": "Monitor Updated",
				"typeRequest": "ACH-VC5-6",
				"amount": 32,
				"price": 6000				
			}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	req, rr = createRequestTest(http.MethodDelete, "/products/1", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}
