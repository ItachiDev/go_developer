package main

import (
	"fmt"
)

func main() {
	// Puntero es la dirección de memoria de las variables
	// El zero value de un puntero es nil
	a := 3
	fmt.Println("Antes de duplicar, a =", a)
	duplicar(&a)
	fmt.Println("Despues de duplicar, a =", a)
}

// A esto se le llama paso por referencia ya que esta recibiendo un puntero
func duplicar(a *int) {
	*a *= 2
	fmt.Println("Dentro de la función duplicar a vale", *a)
}
