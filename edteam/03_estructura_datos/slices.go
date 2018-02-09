package main

import (
	"fmt"
)

func main() {
	// Slices representan una secuencia de tamaño variable de elementos del mismo tipo
	// Los slices tienen un aputador a un array
	// 1.- Apuntador a un array
	// 2.- Tamaño que no es fijo
	// 3.- Capacidad

	var nombres []string
	fmt.Printf("Su tamaño es: %d y su capacidad es %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Daniel")
	fmt.Printf("Su tamaño es: %d y su capacidad es %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Ana")
	fmt.Printf("Su tamaño es: %d y su capacidad es %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Laura")
	fmt.Printf("Su tamaño es: %d y su capacidad es %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Pablo")

	// Obteniendo el tamaño y la capacidad del slice
	fmt.Printf("Su tamaño es: %d y su capacidad es %d\n", len(nombres), cap(nombres))

	// Asignar un valor a un dato que ya existia previamente
	nombres[3] = "Juan"

	// Valores de Slice
	fmt.Println(nombres)

	// Forma más comun de como crear Slices -- funcion make
	// 1.- Tipo de datos
	// 2.- Tamaño Inicial
	// 3.- Capacidad Inicial (Opcional)
	names := make([]string, 0)

	// Crear e inicializar slice
	aplicaciones := []string{
		"Facebook",
		"Twitter",
		"Likedin",
		"Instagram",
		"YouTube",
	}
}
