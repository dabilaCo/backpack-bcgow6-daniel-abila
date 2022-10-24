package main

import (
	"log"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Repaso-Final/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	err := r.Run(":18085")
	if err != nil{
		log.Fatal("error al inciar")
	}
}
