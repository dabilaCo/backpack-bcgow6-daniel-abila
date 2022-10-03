package main

import(
	"fmt"
)

var temperatura int
var humedad int
var presionAtmosferica float32

func main(){
temperatura = 16
humedad = 64
presionAtmosferica = 1024

fmt.Println("Temperatura:", temperatura)
fmt.Println("Humedad:", humedad)
fmt.Println("Presion Atmosférica:", presionAtmosferica)
}