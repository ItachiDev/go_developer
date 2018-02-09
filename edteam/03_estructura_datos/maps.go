package main

import "fmt"

func main() {
	// Mapas: son estructuras de datos clave - valor que permiten tener claves personalizadas
	// Diccionario o Arrays asociativos en otros lenguajes de programación

	// Shortcut para declara mapas
	idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Ingles"
	idiomas["fr"] = "Frances"

	fmt.Println(idiomas)

	// Asignación de valor a mapas
	languages := map[string]string{
		"es": "Español",
		"en": "Ingles",
		"fr": "Frances",
	}
	fmt.Println(languages)

	// Imprimir una sola posición del mapa
	fmt.Println(languages["fr"])

	// Borrar una posición del mapa
	delete(languages, "en")
	fmt.Println(languages)

	// Los mapas devuelven el valor de la clave y un booleano si existe o no
	fmt.Println(languages["pt"])
	if language, ok := languages["es"]; ok {
		fmt.Println("La posición existe y vale:", language)
	} else {
		fmt.Println("La posición no existe")
	}

	numeros := map[int]string{
		1:  "Numero inicial",
		2:  "Segundo numero",
		16: "Tamaño de fuente default",
		-4: "Numero negativo como index",
	}
	fmt.Println(numeros[-4])

}
