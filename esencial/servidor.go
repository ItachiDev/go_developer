package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", saludo)
	http.ListenAndServe(":9696", nil)
}

func saludo(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hola mundo")
}
