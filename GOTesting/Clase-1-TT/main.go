package main

import (
	"fmt"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-1-TT/pkg/operaciones"
)

func main() {
	a, b := 10, 5
	sum, _ := operaciones.Add(a, b)
	fmt.Println(sum)
}
