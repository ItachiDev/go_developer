package main

import (
	"fmt"
)

func main() {
	var arreglo [2]string
	arreglo[0] = "Manzanas"
	arreglo[1] = "Bananas"

	arreglo_dos := [2]string{"Piñas", "Peras"}
	arreglo_dos[1] = "Uvas"

	fmt.Println(arreglo)
	fmt.Println(arreglo_dos)

	// arrays multidimensionales
	var frutas [2][2]string
	// fruta numero uno
	frutas[0][0] = "Manzanas"
	frutas[0][1] = "Golden"
	// fruta numero 2
	frutas[1][0] = "Uvas"
	frutas[1][1] = "Moscatel"
	fmt.Println(frutas)
	fmt.Println(frutas[1])
}
