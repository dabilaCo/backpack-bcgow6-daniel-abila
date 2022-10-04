package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Definimos una pseudobase de datos donde consultaremos la información.
var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

func main() {
	server := gin.Default()
	server.GET("/", Home)
	server.GET("/empleados", SearchEmployee)
	server.Run(":8080")
}

func Home(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pagina Principal \n")
}

func SearchEmployee(ctx *gin.Context) {
	empleado, err := empleados[ctx.Query("id")]
	if err {
		ctx.String(200, "Información del empleado %s, nombre: %s \n", ctx.Query("id"), empleado)
	} else {
		ctx.String(404, "Información del empleado ¡No existe! \n")
	}
}
