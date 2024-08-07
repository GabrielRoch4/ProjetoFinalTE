package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

// ProfessorRoutes configura as rotas relacionadas a professores
func ProfessorRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/professores", controllers.GetProfessores)
	mux.HandleFunc("/professores/", controllers.GetProfessorPorID)
	mux.HandleFunc("/professores/cadastrar", controllers.CreateProfessor)
	mux.HandleFunc("/professores/atualizar", controllers.UpdateProfessor)
	mux.HandleFunc("/professores/deletar", controllers.DeleteProfessor)
}
