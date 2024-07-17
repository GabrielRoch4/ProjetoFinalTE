package main

import (
	"ProjetoFinal/routes"
	"net/http"
)

func main() {

	routes.TurmaRoutes()
	routes.ProfessorRoutes()
	routes.AlunoRoutes()
	routes.AtividadeRoutes()

	_ = http.ListenAndServe(":3333", nil)
}
