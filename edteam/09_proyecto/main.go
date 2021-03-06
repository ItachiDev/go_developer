package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/go_developer/edteam/09_proyecto/migration"
	"github.com/go_developer/edteam/09_proyecto/routes"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzó la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración...")
	}

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := http.Server{
		Addr:    "8080",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecución del programa")
}
