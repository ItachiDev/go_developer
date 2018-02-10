package main

import (
	"fmt"
	// se utiliza "_" cuando no se va a utilizar el paquete explicitamente dentro de la aplicación
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Producto contiene las estructura de un articulo
type Producto struct {
	// GORM crea automaticamente campos:
	// ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@/dev_edesk?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexion a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conecto a la base de datos")

	// Se crea una tabla con las estructura del puntero que le pasamos por parametro
	/* db.CreateTable(&Producto{})
	fmt.Println("Se creo la tabla en la base de datos") */

	// Registra un producto en la tabla
	/*db.Create(&Producto{
		CodigoBarras: "0202020202",
		Precio:       800,
	})
	fmt.Println("Se creo el registro en la base de datos")*/

	// Consultar registro de la base de datos
	var p Producto
	/* db.First(&p)  Obtenemos el primer registro de la tabla*/
	db.First(&p, 2) //Obtenemos el registro con id => 2
	fmt.Println("Precio del producto consultado:", p.Precio)

	// Guardar los cambios que se hagan para la estructuta
	p.Precio = 750
	db.Save(&p)
	fmt.Println("Precio del producto consultado:", p.Precio)
}
