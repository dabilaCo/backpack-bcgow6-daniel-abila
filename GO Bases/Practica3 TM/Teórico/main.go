package main

import "fmt"

func main() {
	/*messageHello := "Hello World :P\n"
	messageBye := "Bye, world :(\n"
	b, err := fmt.Print(messageHello, messageBye)
	if(err != nil){
		panic(err.Error())
	}*/

	//fmt.Print(b)
/*Ejemplo Printf
	flag := true
	flag2 := false
	fmt.Printf("la variable es %v. La otra es %v", flag, flag2)*/

	//number := 12222.123123

	//fmt.Printf("El número resumido es %.3f,", number)

	nombre := "Daniel"
	edad:= 25

	//fmt.Printf("Hola, %s! Tienes %d años.", nombre, edad)

	//Sprint
	message := fmt.Sprint("Hola ", nombre, " tienes ", edad, " años")
	fmt.Println(message)
}
