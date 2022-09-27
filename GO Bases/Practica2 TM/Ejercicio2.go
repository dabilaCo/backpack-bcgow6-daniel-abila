package main

import (
	"errors"
	"fmt"
)

func promedioCalificacion(nums ...float64) (float64, error) {
	var promedio float64
	var suma float64
	var aux float64
	for _, value := range nums {
		if value >= 0 {
			suma += value
			aux++
		} else {
			return 0, errors.New("No podemos tener valores negtivos")
		}

	}

	promedio = suma / aux
	return promedio, nil
}

func main() {
	result, err := promedioCalificacion(2, 4, 6, -3, 10)

	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}
}
