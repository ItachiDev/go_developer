package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	// Un handle o manejador es una interfaz que no respite recibir trabajar con responses y requests
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un Handler</h1>"}
	m2 := &messageHandler{message: "<h1>Hola desde un Handler 2</h1>"}
	m3 := &messageHandler{message: "<h1>Adios desde un Handler </h1>"}

	mux.Handle("/saludar", m1)
	mux.Handle("/saludo", m2)
	mux.Handle("/despedirse", m3)
	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
