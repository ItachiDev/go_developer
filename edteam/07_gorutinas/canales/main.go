package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "go", "git"}
	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	// Se carga la fotocopiadora
	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	// Se deja en la bodega destino
	go func() {
		for {

			// entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora
			fotocopiado <- libro

			// agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {

				// cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("Estado de libros fotocopiados")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
}
