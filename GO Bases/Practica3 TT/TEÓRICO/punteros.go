package main

import "fmt"

func main(){

	//Crear un puntero
	//var puntero *int
	//puntero2 := new(int)

	//Obtener puntero a partir de una variable
	var numero int  = 10

	p3 := &numero

	fmt.Printf("La dirección es: %d\n", numero)
	//operador de resreferenciación  sirve para modificar el valor de una variable desde afuera del scope
	//*p3 = 10
	incrementar(p3)

	fmt.Printf("La dirección es: %d\n", numero)
}

func incrementar(puntero *int){

	*puntero = 20
}

type Coso struct{
	Valor int
}

/*func (c Coso) Actualizar(new int){
	c.Valor 
}*/