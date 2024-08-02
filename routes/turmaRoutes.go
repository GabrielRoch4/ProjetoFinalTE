package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func TurmaRoutes(mux *http.ServeMux) {
	http.HandleFunc("/turmas", controllers.GetTurma)
	http.HandleFunc("/turmas/cadastrar", controllers.CreateTurma)
	http.HandleFunc("/turmas/atualizar", controllers.UpdateTurma)
	http.HandleFunc("/turmas/deletar", controllers.DeleteTurma)
}
