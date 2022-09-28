package main

import (
	"fmt"
	"os"
)

type product struct {
	idProduct int
	price     float64
	amount    int
}

func (producte product) format() string {
	return fmt.Sprintf("ID: %d\n Price: %f\n Amount: %d", producte.idProduct, producte.price, producte.amount)
}

func main() {
	product1 := product{
		idProduct: 123,
		price:     1400,
		amount:    4,
	}

	product2 := product{
		idProduct: 453,
		price:     2000,
		amount:    7,
	}

	os.Create("products.csv")

	f, err := os.OpenFile("products.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte(product1.format())); err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte(product2.format())); err != nil {
		panic(err)
	}

}
