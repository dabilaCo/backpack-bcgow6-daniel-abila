package main

import (
	"fmt"
) //Esto es para poder mostrar en pantalla

//var horas int  //manera 1 de crear una variable

func main() {
	//var horas = 25   //manera 2 de creaar una variable
	horas:= 25   //manera 3 de crear una variable
	const segundos int = 34  //esto es una constante, no se modifica y se debe de inicializar con un valor
	fmt.Println(horas)
	//fmt.Printf("%T", float64(horas))  //esto te indica el tipo de dato de la variable horas - El float64 es como un casteo
	//pero no cambia el tipo de dato de la variable horas, solo en esta consulta
	fmt.Println(segundos)
}
