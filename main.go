package main

import (
	"ProjetoFinal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/professor", handlers.ProfessorHandler)

	_ = http.ListenAndServe(":3333", nil)
}
