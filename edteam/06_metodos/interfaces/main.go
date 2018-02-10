package main

import "go_developer/edteam/06_metodos/interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Cerbero"
	g.Nombre = "Matu"

	//AdoptarPerro(p)
	//AdoptarGato(g)

	// Para poder implemantar una interfaz el metodo no puede recibir un puntero
	AdoptarMascota(p)
	AdoptarMascota(g)
}

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}

/*
	func AdoptarPerro(p animales.Perro) {
		p.Comunicarse()
	}

	func AdoptarGato(g animales.Gato) {
		g.Comunicarse()
	}
*/
