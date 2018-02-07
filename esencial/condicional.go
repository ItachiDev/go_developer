package main

import (
	"fmt"
)

func main() {
	puntos := 101

	if puntos < 10 {
		fmt.Println("Puntos Incorrectos!")
	} else if puntos == 100 {
		fmt.Println("puntos son Correctos")
	} else {
		fmt.Println("Tus puntos son ", puntos)
	}
}
