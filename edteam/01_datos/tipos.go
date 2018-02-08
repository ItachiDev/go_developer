package main

import "fmt"

func main() {
	// Tipo de datos en Go
	// En Go no podemos hcer operaciones aritmeticas con tipos de datos diferentes

	// int es el alias para el compilador de int64
	var a int
	var b int8

	a = 123456
	b = 5

	// Go nos permite hacer casting de variables
	// casting es hacer el cambio temporal de una variable para hacer una operación especifica
	c := a + int(b)
	fmt.Println(c)

	// Interpolación de una variable en una cadena de texto
	name := "Pedro"
	fmt.Printf("hola %s cómo estas?\nSon las %d p.m\n", name, 5)

	// Obtener el tipo de valor de la sintaxis de Go
	fmt.Printf("El tipo de dato de la variable es: %T", b)

	// Valores 0 de las variables
	// Todos los tipo de datos tienen su valor inicial llamado zero value
	var vacio string
	var cero int
	var falso bool
	fmt.Println(vacio, cero, falso)
}
