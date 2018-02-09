package main

import (
	"errors"
	"fmt"
)

func main() {
	// Los errores son un tipo de dato
	// El zero value de un tipo error es nil

	err := errors.New("Mi mensaje de error")
	fmt.Printf("%T\n", err)

	resultado, errores := division(100, 2)
	if err != nil {
		fmt.Println("Error:", errores)
		return
	}
	fmt.Println(resultado, errores)
}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir entre 0")
	} else {
		resultado = dividendo / divisor
	}
	return
}
