package main

import "fmt"

func main() {
	// Todo lo que mandemos a defer se ejecutara al final
	// Si en algun punto llega a encontrar un return se ejecutaran los defer
	// Si la aplicación entra en panic... Se ejecutará la cola de defer
	// Los defer se ejecutan en orden inverso

	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()

	fmt.Println("Consultamos información, set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDeDatos() {
	fmt.Println("Cierra el set de datos")
}
