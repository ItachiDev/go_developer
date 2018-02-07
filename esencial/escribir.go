package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contenido := []byte("Este es un dato almacenado desde GO")
	datos := ioutil.WriteFile("archivo.txt", contenido, 0755)

	mostrarError(datos)
	fmt.Println("La información ha sido escrita")
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
