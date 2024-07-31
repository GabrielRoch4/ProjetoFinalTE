package main

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/routes"
	"net/http"
)

func main() {

	routes.TurmaRoutes()
	routes.ProfessorRoutes()
	routes.AlunoRoutes()
	routes.AtividadeRoutes()

	database.DatabaseConnection()

	_ = http.ListenAndServe(":3333", nil)
}
