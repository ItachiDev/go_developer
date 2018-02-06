package main

import "fmt"

func main() {

	// Un slice tiene una cantidad dinamica de elementos
	var frutas = []string{"Manzanas", "Uvas", "Bananas", "Peras"}
	fmt.Println(frutas)

	// Los slices son instancias de un array donde no debemos de definir un tamaño de elementos
	// Tamaño de items del slice
	fmt.Println(len(frutas))

	// Agregar elementos de manera dinamica
	frutas = append(frutas, "Sandias")
	fmt.Println(frutas)

	// Vizualizar rangos dentro de un slice
	fmt.Println(frutas[1:3])
}
