package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func TurmaRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/turmas", controllers.GetTurma)
	mux.HandleFunc("/turmas/", controllers.GetTurmaPorID)
	mux.HandleFunc("/turmas/cadastrar", controllers.CreateTurma)
	mux.HandleFunc("/turmas/atualizar", controllers.UpdateTurma)
	mux.HandleFunc("/turmas/deletar", controllers.DeleteTurma)
}
