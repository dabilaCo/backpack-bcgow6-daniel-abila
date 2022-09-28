package main

import (
	"errors"
	"fmt"
)

func main() {

	var salary int
	salary = 150000

	if salary < 150000 {
		fmt.Println(errors.New("el salario ingresado no supera el mínimo de 150.000."))
		return
	}

	fmt.Println("el salario ingresado supera el mínimo de 150.000, debe pagar impuestos.")
}
