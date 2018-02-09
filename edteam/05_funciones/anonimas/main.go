package main

import (
	"fmt"
)

// Una funcion anonima es un tipo de dato en Go
// Una funcion anonima mantiene disponible el valor de sus variables para la siguiente ejecución
// Puede ser asignada a una variable

func main() {

	anonima := func() {
		fmt.Println("Me imprimo estando en una variable")
	}
	anonima()

	secuencs := secuencia()
	fmt.Println(secuencs())
	fmt.Println(secuencs())
	fmt.Println(secuencs())
	fmt.Println(secuencs())

	secuencsDos := secuencia()
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
	fmt.Println(secuencsDos())
}

// Pueden existir funciones que retornen funciones
func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
