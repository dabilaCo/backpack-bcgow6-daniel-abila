package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id            int 
	Nombre        string
	Apellido      string
	Email         string
	Edad          int
	Height        int
	Activo        bool
	FechaCreacion string
}


func generarUsuarios() []Usuario {
	usuarios := []Usuario{
		{1, "Juan", "Perez", "juanandres.perez@mercadolibre.com.co", 19, 176, true, "19-09-2022"},
		{2, "John", "Doe", "johndoe@enterprises.co", 50, 160, false, "14-04-2021"},
		{3, "Jane", "Doe", "jane.doe@noidentity.wal.es", 45, 167, true, "01-12-2019"},
	}
	return usuarios
}

func isNil(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAll(c *gin.Context) {
	//creo una lista de usuarios
	usuarios := generarUsuarios()
	c.JSON(http.StatusOK, gin.H{
		"total":      len(usuarios),
		"resultados": usuarios,
	})

}

func main() {

	//creo un router con gin
	router := gin.Default()

	/*router.GET("/hello-daniel", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Daniel",
		})
	})*/

	router.GET("/usuarios", GetAll)
	router.GET("/usuarios/:id", getUser)
	//corremos nuestro servidor sobre el puerto 8080
	router.Run(":8080")

}

func getUser(c *gin.Context){
	usuarios := generarUsuarios()
	usuario, err := usuarios[c.Param("id")]
}


