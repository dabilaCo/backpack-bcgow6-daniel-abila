package main

import "fmt"

// creo el struct del error
type myCustomError struct {
}

// manejo el error personalizado
func (e *myCustomError) Error() string {
	return fmt.Sprintf("el salario ingresado no supera el mínimo de 150.000.")
}

// acá valido que el salario ingresado llegue a 150.000
func checkSalary(salary int) (err error) {
	if salary < 150000 {
		err = &myCustomError{}
	}
	return
}

func main() {
	var salary int

	salary = 75000
	err := checkSalary(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("el salario ingresado supera el mínimo de 150.000, debe pagar impuestos.")
	}

}
