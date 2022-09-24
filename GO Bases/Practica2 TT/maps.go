package main

import "fmt"

func main() {
	var mapa = map[string]int{"Nico": 23, "Daniel": 25}
	mapa["Hola"] = 10

	x, ok := mapa["Hola"]

	fmt.Println(x, ok)

	
}
