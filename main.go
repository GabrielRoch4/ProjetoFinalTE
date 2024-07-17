package main

import (
	"ProjetoFinal/routes"
	"net/http"
)

func main() {

	routes.TurmaRoutes()
	routes.ProfessorRoutes()

	_ = http.ListenAndServe(":3333", nil)
}
