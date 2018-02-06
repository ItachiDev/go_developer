package main

import "fmt"

func main() {
	// Go es un lenguaje de programación tipado
	var cadena string
	cadena = "Development"
	fmt.Println(cadena)

	var entero int
	entero = 100
	fmt.Println(entero)

	var flotante float32
	flotante = 0.10119228282
	fmt.Println(flotante)

	// Asignación dinamica en Go
	dinamico := "valor dinamico"
	fmt.Println(dinamico)

	// Constantes en Go
	const estatico bool = true
	fmt.Println(estatico)

	// Operaciones Aritmeticas (+, -, *, /)
	var numero int
	numero = 10
	fmt.Println(numero + 10)

	// Difinir tipos de datos personalizados
	var miCafe = Cafe{nombre: "espresso", precio: 5.22, azucar: false, leche: 0}
	fmt.Println(miCafe)
}

type Cafe struct {
	nombre string
	precio float64
	azucar bool
	leche  int
}
