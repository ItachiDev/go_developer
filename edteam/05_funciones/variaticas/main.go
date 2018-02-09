package main

import "fmt"

func main() {
	// Una funcion variatica recibe un numero variable de parametros del mismo tipo
	saludarVarios(25, "Cristhofer", "Arturo", "Gerardo", "Daniel", "Ana")
}

// Las funciones variaticas tiene varias restricciones, solo podemos tener un parametro variatico
// El parametro variatico debe estar al final de función
func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)
	for _, value := range nombres {
		fmt.Println("Hola como estas", value, "tiene", edad, "en promedio")
	}
}
