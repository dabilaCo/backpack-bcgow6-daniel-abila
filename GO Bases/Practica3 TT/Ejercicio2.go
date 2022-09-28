package main

import "fmt"

// producto
type Producto struct {
	Nombre   string
	precio   float64
	cantidad int
}

// Usuario
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

// crear u nuevo producto
func newProduct(nombre *string, precio *float64) *Producto {
	return &Producto{Nombre: *nombre, precio: *precio}
}

// Agregar un producto a un Usuario
func (u *Usuario) addProduct(product *Producto, cantidad *int) {
	product.cantidad = *cantidad
	u.Productos = append(u.Productos, *product)
}

// Eliminar un producto de un Usuario
func (u *Usuario) deleteProduct() {
	u.Productos = []Producto{}
}

func main() {

	//Creo nuevo producto
	var (
		nombre   string  = "Casancrem"
		precio   float64 = 140
		cantidad int     = 3
	)
	producto1 := Producto{Nombre: "Stanley", precio: 3000, cantidad: 4}

	product1 := newProduct(&producto1.Nombre, &producto1.precio)
	product2 := newProduct(&nombre, &precio)

	//Creo un usuario
	usuario := &Usuario{
		Nombre:   "Daniel",
		Apellido: "Abila",
		Correo:   "danielabila03@gmail.com",
	}

	usuario.addProduct(product1, &producto1.cantidad)
	usuario.addProduct(product2, &cantidad)

	fmt.Println("Usuario - ", usuario.Nombre, usuario.Apellido)
	fmt.Println("Correo - ", usuario.Correo)
	for _, value := range usuario.Productos {
		fmt.Printf("Producto - %s - Precio %.2f Cantidad %d - Total $%.2f\n", value.Nombre, value.precio,
			value.cantidad, (float64(value.cantidad) * value.precio))
	}

	fmt.Println()
	fmt.Println("Eliminando productos . . .")
	usuario.deleteProduct()

	fmt.Println("Cantidad de productos después de la eliminación", len(usuario.Productos))

}
