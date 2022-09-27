package main

import (
	"errors"
)

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func minFunc(students ...int) (float64, error) {
	min := students[0]

	for _, value := range students {
		if value < 0 {
			return 0, errors.New("Valor negativo")
		}
		if value < min {
			min = value
		}
	}
	return float64(min), nil
}

func maxFunc(students ...int) (float64, error) {
	max := students[0]

	for _, value := range students {
		if value < 0 {
			return 0, errors.New("Valor negativo")
		}
		if value > max {
			max = value
		}
	}
	return float64(max), nil
}

func averageFunc(students ...int) (float64, error) {
	var suma int
	var aux int
	var promedio float64
	for _, value := range students {
		if value < 0 {
			return 0, errors.New("Valor negativo")
		}
		if value >= 0 {
			suma += value
			aux++
		} else {
			return 0, errors.New("Valor negativo")
		}

	}

	promedio = float64(suma / aux)
	return float64(promedio), nil
}

func operation(op string) (func(values ...int) (float64, error), error) {
	switch op {
	case "minimum":
		return minFunc, nil
	case "maximum":
		return maxFunc, nil
	case "average":
		return averageFunc, nil
	default:
		return nil, errors.New("Invalid argument")

	}
}

func main() {

}
