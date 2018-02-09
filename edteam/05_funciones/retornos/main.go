package main

import (
	"fmt"
)

func main() {
	var n1, n2 int8
	n1 = 15
	n2 = 6

	// Función para sumar dos valores
	result := suma(n1, n2)
	fmt.Println(result)

	// Función para obtener el tipo de edad
	var edad uint8
	edad = 3
	fmt.Println(tipoEdad(edad))

	// Obtener los valores maximos y minimos
	n := []int8{23, 94, 34, 83, 65, 100, 47, 99, 1, 12, 40}
	max, min := maxMin(n)
	fmt.Println(max, min)

}

func suma(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "niñez"
	case edad < 18:
		tipo = "adolescencia"
	default:
		tipo = "adulto"
	}

	return tipo
}

// Go tiene la posibilidad de retornar multiples valores
func maxMin(numeros []int8) (max int8, min int8) {
	for _, value := range numeros {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return
}
