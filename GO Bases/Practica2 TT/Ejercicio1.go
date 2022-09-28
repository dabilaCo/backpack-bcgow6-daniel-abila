package main

import "fmt"

type Students struct {
	firstName string
	lastName string
	id int
	date string
}

//esto es un método
func (student Students) detail(){
	fmt.Println("Nombre:",student.firstName)
	fmt.Println("Apellido:",student.lastName)
	fmt.Println("DNI:",student.id)
	fmt.Println("Fecha ingreso:",student.date)
}

func main(){
	student := Students{"Daniel", "Abila", 47418557, "27 de Septiembre"}

	student.detail()
}