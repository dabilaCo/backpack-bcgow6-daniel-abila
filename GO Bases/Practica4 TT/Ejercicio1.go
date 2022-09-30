package main

import (
	"fmt"
	"os"
)


func main(){
	
	defer func(){
		fmt.Println("ejecución finalizada")
	}()

	data, err := os.ReadFile("./customer.txt")
	if err != nil{
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	os.Stdout.Write(data)

	

	

}