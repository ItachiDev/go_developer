package main

import "fmt"

func main() {
	var a uint8
	a = 5

	// En Go no se requiere la palabra reservada break para la estructuta switch
	switch a {
	case 1:
		fmt.Println("El valor es igual a uno")
	case 2:
		fmt.Println("El valor es igual a dos")
	default:
		fmt.Println("No es ninguno de los casos previos")
	}

	// fallthrough en structuras switch continua evaluando hasta encontrar una intrucción
	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("Estas entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un día valido")
	}

	// switch nos permit declarar la variable dentro de la evaluación del mismo
	switch a := 3; {
	case a >= 0 && a <= 5:
		fmt.Println("Estas entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estás en fin de semana")
	default:
		fmt.Println("No es un día valido")
	}
}
