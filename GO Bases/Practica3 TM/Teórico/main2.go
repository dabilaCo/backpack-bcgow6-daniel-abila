package main

import (
	"fmt"
	"os"
)

func main(){

	/*name, ok := os.LookupEnv("HOME")
	if !ok{
		panic("No se encuentra la variable HOME")
	}

	//Setenv
	/*err := os.Setenv("NAME", "Daniel")
	if err!= nil{
		panic(err)
	}*/

	//Getenv
	name := os.Getenv("HOME") 
	fmt.Println("RESULTADO: ",name)
}