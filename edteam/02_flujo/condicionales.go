package main

import "fmt"

func main() {

	// Condicionales en Go
	// || Or
	// && And
	// ! Not
	// En este caso la variable solo vive en el scope de la condicion

	if isValid := true; isValid {
		fmt.Println("Esto está en el bloque de verdadero")
		fmt.Printf("%T\n", isValid)
	} else {
		fmt.Println("Esto esta en el bloque falso")
	}

	edad, nombre := 10, "Cristhofer"
	if edad < 14 {
		fmt.Println(nombre, "es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "es un adolescente")
	} else {
		fmt.Println(nombre, "es un adulto")
	}

	fmt.Println("Fin del programa")
}
