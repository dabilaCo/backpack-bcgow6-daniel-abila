package main

import "fmt"

type Matrix struct {
	Values      []float64
	Height      int
	Width       int
	IsQuadratic bool
	MaxValue    int
}

func (m *Matrix) Set(values ...float64) { m.Values = values }

func (m Matrix) Print() {
	for line := 0; line < m.Height; line++ {
		for value := 0; value < m.Width; value++ {
			fmt.Printf("%f\t", m.Values[line*m.Width+value])
		}
		fmt.Printf("\n")
	}
}

func main() {
	m := Matrix{Height: 3, Width: 3, IsQuadratic: true}

	m.Set(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)

	m.Print()
}
