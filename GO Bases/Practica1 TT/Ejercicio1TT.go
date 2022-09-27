package main

import (
	"fmt"
)

func main() {
	var palabra = "Daniel"
	
	char := []rune(palabra)

	fmt.Println("La palabra", palabra, "tiene", len(palabra), "letras")

	fmt.Println("")
	fmt.Println("Sus letras son:")

	for i := 0; i < len(char); i++{
		letra := string(char[i])
		fmt.Println(letra)
	}

	
}