package main

import (
	"fmt"
	"time"
)

func main() {
	tiempo := time.Now()
	switch diaDeHoy := tiempo.Weekday(); diaDeHoy {
	case 0:
		fmt.Println("Descanso de Domingo")
	case 1:
		fmt.Println("Odio los Lunes")
	case 2:
		fmt.Println("El martes no lo odio tanto")
	default:
		fmt.Println("Tango que trabajar")
	}

}
