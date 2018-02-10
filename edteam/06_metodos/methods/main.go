package main

import "fmt"

// estructura Persona
type Persona struct {
	nombre string
	edad   int8
}

type Numero int

func (person Persona) saludar() {
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse() {
	fmt.Println("Hola yo soy un numero")
}

// Si quiero recibir algo del tipo de dato debemos de recibir un puntero
func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("No es validad la cantidad edad")
	}
}

func main() {
	// En podemos asociar un metodo al cualquier tipo de dato
	// En Go no es necesario asignar estructura a un metodo solamente
	var persona Persona
	persona.saludar()

	persona.asignarEdad(-35)
	fmt.Println(persona)

	var numero Numero
	numero.presentarse()
}
