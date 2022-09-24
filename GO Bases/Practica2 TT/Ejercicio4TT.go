package main

import "fmt"

func main(){

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var contador int
	//Edad  de Benjamin
	fmt.Println("Benjamin tiene", employees["Benjamin"], "años.")

	//Cuantos son mayores de 21
	for key, element := range employees{
		if(element > 21){
			contador++
		fmt.Println(key, element)
		}
	}

	fmt.Println("Hay", contador, "empleados mayores de 21")

	//Agregar empleado
	employees["Federico"] = 25

	//Eliminar
	delete(employees,"Pedro")





}