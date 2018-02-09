package main

import (
	"fmt"
)

func main() {
	// El for continuo solo contiene la expresión de condición
	a := 5

	for a > 0 {
		a--
		fmt.Println(a)
	}

	// Ciclo infinito "forever"
	// No tiene ninguna condicion que lo detenga
	for {
		fmt.Println("Hola, nunca voy a detener este proceso")
	}
}
