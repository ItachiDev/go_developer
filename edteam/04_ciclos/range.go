package main

import (
	"fmt"
)

func main() {
	numeros := []string{"uno", "dos", "tres", "cuatro"}

	// Podemos evitar usar una variable si dentro de la instrucción del ciclo
	// utilizamos "_" para ignorar la variable.

	// El valor en un for range puede omitirse, el indice ¡no!
	for index, value := range numeros {
		fmt.Println(index, value)
	}

	nombres := map[string]string{"es": "Español", "co": "Colombia", "br": "Brasil"}
	for index, value := range nombres {
		fmt.Println(index, value)
	}

	// Podemos iterar strings
	// Para go una letra es un numero (byte) por lo cual se castea el dato
	frase := "Hola mundo"
	for index, value := range frase {
		fmt.Println(index, string(value))
	}

	// Pasar directamente el valor al for range
	for _, numero := range []int{15, 36, 24, 85} {
		fmt.Println(numero)
	}
}
