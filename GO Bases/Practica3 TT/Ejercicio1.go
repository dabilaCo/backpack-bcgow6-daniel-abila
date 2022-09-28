package main

import "fmt"

//Estructura
type Persona struct{
	Nombre string
	Apellido string
	Edad int
	Correo string
	Password string
}

//cambiarNombre
func (p *Persona) cambiarNombre(nombre string, apellido string){
	p.Nombre = nombre
	p.Apellido = apellido	
}

//cambiarEdad
func (p *Persona) cambiarEdad(edad int){
	p.Edad = edad
}

//cambiarCorreo
func (p *Persona) cambiarCorreo(correo string){
	p.Correo = correo
}

//cambiarContraseña
func (p *Persona) cambiarContrasena(nuevaPass string){
	p.Password = nuevaPass
}

func main(){

	persona := Persona{Nombre: "Chiara", Apellido: "Pozzi", Edad: 25, Correo: "chiarapozzig@gmail.com", Password: "chiara14"}

	fmt.Println("Persona sin modificar:", persona)

	persona.cambiarContrasena("chiara1414")

	fmt.Println("Persona modificando Contraseña:", persona)

	persona.cambiarCorreo("chiarapozzig@hotmail.com")
	persona.cambiarNombre("Alberto", "Ganduglia")

	fmt.Println(persona)

}