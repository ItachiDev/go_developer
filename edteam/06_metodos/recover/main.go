package main

import "fmt"

func main() {
	// recover recupera la aplicacion cuando entra en panico
	// recover se utilizan dentro de la funciones defer
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuoerado en f:", r)
		}
	}()
	fmt.Println("Llamando a g.")
	g(5)
	fmt.Println("Retorna normalmente g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Entre en panico")
		panic("numero no puede ser mayor que 3")
	}
	defer fmt.Println("Defer de la función g")
	fmt.Println("Imprimiendo en g", i)
}
