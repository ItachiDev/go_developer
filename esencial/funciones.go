package main

import "fmt"

func main() {
	fmt.Println(multiplicar(20, 100))
	fmt.Println(saludar())
	fmt.Println(ganado(1))
	alumnos("Hugo", "Fracisco", "Erick")
}

// función que devolvera un dato de tipo entero
func multiplicar(numero1, numero2 int) int {
	return numero1 * numero2
}

// argumentos y retorno en las funciones
func saludar() (a string, b int) {
	a = "Hola"
	b = 77
	return
}

// closures en Go
func ganado(numero int) (int, string) {
	vacas := func() int {
		return numero * 10
	}
	return vacas(), "vacas"
}

// uso de argumentos, podemos recibir 1 ó n cantidad de argumentos
func alumnos(alumno ...string) {
	for _, valor := range alumno {
		fmt.Println(valor)
	}
}
