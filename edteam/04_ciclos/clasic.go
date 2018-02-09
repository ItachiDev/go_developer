package main

import (
	"fmt"
)

func main() {
	// Clico for clasico
	// 1.- Sentencia de Inicio ";"
	// 2.- Expresión de condición ";"
	// 3.- Instrucción post ejecución

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 4; i >= 0; i-- {
		// Permite saltar el resto de instrucciones de la iteración actual
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// Detener el ciclo con la intrucción break
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("i vale 2 y se rompe el ciclo")
			break
		}
		fmt.Println(i)
	}

	// Generar una matriz con ciclos for
	matriz := [3][3]int{}
	valor := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}

	fmt.Println(matriz)

	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}
	}
}
