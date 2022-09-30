package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("./panic.go")
	if err != nil{
		panic(err)
	}

	/*var arr []int
	a := arr[4]*/

	//var a *int

	//b := 10

	//a = &b


	file, _ := os.Open("./panic.go")

	defer file.Close()

	
	//DEFER
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("Eu, ocurrió un panic", err)
		}
		fmt.Println("Funciono correctamente")
	}()

	//situación panic
	//ch := make(chan int)
	//fmt.Println(<-ch)	

	retorno, err := FuncionImportante(0)
	if err != nil{
		panic(err)
	}
	fmt.Println("Retorno:", retorno)

	
	panic(10)
	//fmt.Println("Estoy despues del panic")
}

func FuncionImportante(estado int) (string, error){
	if estado == 0{
		return "", errors.New("No hay estado")
	}

	return "Si hay estado", nil
}
