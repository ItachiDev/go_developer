package saludar

import (
	"fmt"
)

// La inicial mayuscula significa que sera una función exportada
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}
