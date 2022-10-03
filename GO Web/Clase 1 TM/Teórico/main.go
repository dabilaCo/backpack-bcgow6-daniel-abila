package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //engine de gin

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "bienvenidos a go web :) \n",
		})
	})

	router.Run()

}
