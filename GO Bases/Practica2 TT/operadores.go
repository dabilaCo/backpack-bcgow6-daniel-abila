package main

import "fmt"

func main(){
	x, y := 10, 20
	//x++
	//y++
	x *= y //esto es lo mismo que hacer x = x * y
	//fmt.Println(x)
	fmt.Println(x > y)
}
