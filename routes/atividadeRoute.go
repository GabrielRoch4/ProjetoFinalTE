package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AtividadeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/turmas/atividades", controllers.GetAtividade)
	mux.HandleFunc("/turmas/atividades/", controllers.GetAtividadePorID)
	mux.HandleFunc("/turmas/atividades/cadastrar", controllers.CreateAtividade)
	mux.HandleFunc("/turmas/atividades/atualizar", controllers.UpdateAtividade)
	mux.HandleFunc("/turmas/atividades/deletar", controllers.DeleteAtividade)
}
