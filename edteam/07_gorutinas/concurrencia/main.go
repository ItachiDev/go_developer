package main

import (
	"fmt"
	"time"
)

func main() {
	var h string

	// Se ejecuta el proceso MostrarNumeros como una Go rutina (proceso independiente)
	go MostrarNumeros()
	fmt.Print("Escribe algo en consola ")
	fmt.Scan(&h)
	fmt.Println("Escribiste:", h)
}

func MostrarNumeros() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
