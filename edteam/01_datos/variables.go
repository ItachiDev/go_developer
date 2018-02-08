package main

import (
	"fmt"
)

func main() {
	/**
	 * para especificar una variable utilizamos:
	 * palabra reservada var
	 * nombre de la variable
	 * tipo de dato
	 * asigno el valor
	 */

	var nombre, apellido string
	nombre = "Cristhofer"
	apellido = "Ordaz"

	fmt.Println(nombre, apellido)

	/**
	 * go tambien puede asignar el tipo de dato dinamicamente con ":="
	 * el operador puede reasignar el valor siempre y cuando asigne el valor a una nueva variable
	 */

	producto, marca := "Inspiron 13 7600", "Dell"
	producto, modelo := "Inspiron 13", "7600"

	fmt.Println(marca, producto, modelo)
}
