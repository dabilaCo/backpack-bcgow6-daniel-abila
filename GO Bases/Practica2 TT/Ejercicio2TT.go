package main

import "fmt"

func main(){
	var edad int 
	var empleado bool
	var antiguedadEnMeses int
	var sueldo int

	
	fmt.Println("Ingrese su edad:")	
	fmt.Scanln(&edad)

	if(edad >= 18){
		fmt.Println("¿Actualmente tiene trabajo? (ingrese true si tiene empleo o false en caso contrario)")
		fmt.Scanln(&empleado)

		if(empleado){
			fmt.Println("Indique su antiguedad en meses:")
			fmt.Scanln(&antiguedadEnMeses)

			if(antiguedadEnMeses >= 12){
				fmt.Println("Ingrese su sueldo:")
				fmt.Scanln(&sueldo)

				if(sueldo > 100000){
					fmt.Println("Usted puede acceder a un préstamo sin interés.")
				}else{
					fmt.Println("Usted puede acceder a un préstamo.")
				}
			}else{
				fmt.Println("Debe de tener una antiguedad de al menos 1 año para acceder a un préstamo.")
			}
		}else{
			fmt.Println("Debe tener empleo para poder acceder a un préstamo.")
		}
	}else{
		fmt.Println("Para acceder a un préstamo debe ser mayor de edad.")
	}

}