package main

import "fmt"

func impuestoSalario(sueldo float64) float64 {
	var impuesto float64

	if sueldo > 50000 && sueldo < 150000 {
		impuesto = sueldo * 0.17
	}

	if sueldo >= 150000 {
		impuesto = sueldo * 0.27
	}

	return impuesto
}

func main() {
	var resultadoImpuesto = impuestoSalario(510000)
	fmt.Println("Se descuenta por impuestos:",resultadoImpuesto)
}
