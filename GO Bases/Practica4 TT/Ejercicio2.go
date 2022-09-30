package main

import (
	"fmt"
	"math/rand"
	"os"
)

//creamos estructura de cliente

type Cliente struct {
	Legajo         int
	NombreApellido string
	DNI            int
	numeroTelefono string
	Domicilio      string
}

func generarIdLegajo() (int, error) {
	//rand es una funcion que te devuelve un numero random
	var id int = rand.Int()
	return id, nil
}

func verificarCliente(id int) {
	//hago un defer porque en la letra pide que se recupere y siga ejecutando luego del error, meto un recover
	defer func() {
		err := recover() //recupera del panic para que siga la ejecución

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("No quedaron archivos abiertos")
	}()

	//leemos el txt que nos menciona la letra
	read, err := os.ReadFile("./customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado.")
	}
	os.Stdout.Write(read)
}

func agregarCliente(legajo, dni int, nombreApellido, domicilio, telefono string) (*Cliente, error) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("se detectaron varios errores en tiempo de ejecución.")
		}
	}()

	if legajo == 0 {
		panic("Legajo no puede ser 0")
	}

	if dni == 0 {
		panic("DNI no puede ser 0")
	}

	if nombreApellido == "" {
		panic("Se requiere el nombre y apellido")
	}

	if domicilio == "" {
		panic("El domicilio es requerido")
	}

	if telefono == "" {
		panic("El telefono es requerido")
	}

	return &Cliente{Legajo: legajo, DNI: dni, NombreApellido: nombreApellido, Domicilio: domicilio, numeroTelefono: telefono}, nil
}

func main() {

	//creo un cliente
	var (
		dni            int    = 47418557
		nombreApellido string = "Daniel Abila"
		telefono       string = "099051993"
		domicilio      string = "Av. Uruguay 1598"
	)

	//En la letra nos pide que el legajo sea generado por separado
	legajoID, err := generarIdLegajo()
	if err != nil {
		panic(err)
	}

	verificarCliente(legajoID)
	agregarCliente(legajoID, dni, nombreApellido, domicilio, telefono)

	fmt.Println("Fin de la ejecución.")

}
