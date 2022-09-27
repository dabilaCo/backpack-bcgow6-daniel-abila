package main

import (
	"errors"
	"fmt"
)

func calculadorSalario(category string, minutesOfWork float64) (float64, error) {

	if minutesOfWork <= 0 {
		return 0, errors.New("si no llegó a un minuto de trabajo, no tiene nada para cobrar.")
	}
	switch category {
	case "a", "A":
		horas := minutesOfWork / 60
		salarioPorHora := 3000 * horas
		salarioBono := salarioPorHora * 0.5
		salario := salarioPorHora + salarioBono
		return float64(salario), nil
	case "b", "B":
		horas := minutesOfWork / 60
		salarioPorHora := 1500 * horas
		salarioBono := salarioPorHora * 0.2
		salario := salarioPorHora + salarioBono
		return float64(salario), nil
	case "c", "C":
		horas := minutesOfWork / 60
		salarioPorHora := 1000 * horas
		return float64(salarioPorHora), nil
	}
	return 0, errors.New("No ingresó una categoría válida, indique A, B ó C.")
}

func main() {
	result, err := calculadorSalario("A", 7000)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Su salario a cobrar es de:",result)
	}

	fmt.Println()
}
