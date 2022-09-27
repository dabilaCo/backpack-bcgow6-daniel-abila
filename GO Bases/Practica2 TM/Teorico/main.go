package main

import "fmt"

/*func inspeccionarVariable(numero int){
	if numero > 0 {
		fmt.Println("El número es positivo")
	} else if numero < 0{
		fmt.Println("El número es negativo")
	} else{
		fmt.Println("El número es cero")
	}
}*/

/*func suma(num1, num2 float64)float64{
	result := num1 + num2
	return result
}*/

const (
	OperadorSuma           = "+"
	OperadorResta          = "-"
	OperadorMultiplicacion = "*"
	OperacionDivision      = "/"
)

func suma(a, b float64) float64 {
	return a + b
}

func resta(a, b float64) float64 {
	return a - b
}

func multiplicacion(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	if b != 0 {
		return a / b
	} else {
		return 0
	}
}

func calculate(operador string) func(float64, float64)float64 {

	switch operador {
	case OperadorSuma:
		return suma
	case OperadorResta:
		return resta
	case OperadorMultiplicacion:
		return multiplicacion
	case OperacionDivision:
		return division
	}
	return nil
}
func main() {
	operation := calculate(OperadorResta)
	
	fmt.Println(operation(10, 2))
}

