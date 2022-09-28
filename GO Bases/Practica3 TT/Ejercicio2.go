package main

import "fmt"

//producto
type Producto struct{
	Nombre string
	precio float64
	cantidad int
}

//Usuario
type Usuario struct{
	Nombre string
	Apellido string
	Correo string
	Productos []Producto
}



func main(){
	 	
	producto1 := Producto{Nombre: "Stanley", precio: 3000, cantidad: 4}
	producto := []Producto{producto1}
	Usuario1 := Usuario{Nombre: "Daniel", Apellido: "Abila", Correo: "danielabila03@gmail.com", Productos:producto }

	fmt.Println(Usuario1)

	

}