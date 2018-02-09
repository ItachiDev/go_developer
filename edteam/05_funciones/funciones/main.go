package main

import (
	"fmt"
)

// La llave de la apertura debe ir siempre en la misma linea del nombre.
func main() {
	saludar("Cristhofer", 27)
}

// Los nombres de las funciones dentro de un mismo paquete deben ser unicos
// Funciones permiten que les pasemos parametros

// return sirve para retornar un valor o detener las instrucciones que continuan debajo la palabra
func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s, tu edad es %d años\n", nombre, edad)
	if edad > 30 {
		return
	}

	fmt.Println("Eres joven aún")
}
