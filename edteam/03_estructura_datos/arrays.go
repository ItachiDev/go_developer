package main

import (
	"fmt"
)

func main() {

	// Arrays
	// Un array es una estructura de datos que nos permite almacenar más
	// de un valor en una sola variable.
	// El zero value de un arreglo es el zero value del tipo de dato que esta almacenando

	// Todos los valores debe ser del mismo tipo
	// Tiene un tamaño fijo
	var nombres [3]string
	nombres[0] = "Cristhofer"
	nombres[1] = "Arturo"
	nombres[2] = "Gerardo"

	fmt.Println(nombres)

	// Declaración del array con shortcut
	names := [3]string{"Daniel", "Ana", "Laura"}
	fmt.Println(names)

	// Obtener tamaño de un arreglo
	size := len(names)
	fmt.Println("El tamaño del arreglo es:", size)

	// Obtener valores de nuestros arreglos
	nombre := names[1]
	fmt.Println(nombre)

	// Obtener una parte del arreglo
	fmt.Println(names[1:3])
}
