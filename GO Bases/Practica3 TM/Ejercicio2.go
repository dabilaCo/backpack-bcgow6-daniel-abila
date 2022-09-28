package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	data, err := os.ReadFile("./products.csv")
	if err != nil{
		panic(err)
	}

	fdata := strings.Split(string(data), "\n")
	for _, line := range fdata {
		if len(line) > 0{
			fline := strings.Split(line, ";")
			fmt.Printf("%s\n", fline[0])
		}
	}
}