package main

import (
	"fmt"
)

// Persona es una estructura
type Persona struct {
	nombre string
	edad   uint8
}

func main() {
	// Estructuras son una secuencia de elementos nombrados que se llaman campos
	// cada campo tiene nombre o tipo.
	// Una estructura es el equivalente a las Clases de otros lenguajes programación

	var person Persona
	person.nombre = "Cristhofer"
	person.edad = 27

	fmt.Println(person)
	fmt.Println(person.nombre)

	// shorthand
	persona := Persona{
		nombre: "Pablo",
		edad:   26,
	}
	fmt.Println(persona)

}
