package main

import "fmt"


func main(){
	var salary int
	salary = 150000

	err := fmt.Errorf("%v", "el salario ingresado no supera el mínimo de 150.000.")

	if salary < 150000 {
		
		fmt.Println(err)
		return
	}else{
		fmt.Println("el salario ingresado supera el mínimo de 150.000, debe pagar impuestos.")
	}

	
}