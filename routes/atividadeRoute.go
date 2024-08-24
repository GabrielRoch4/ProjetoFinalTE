package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AtividadeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/atividades", controllers.GetAtividades)
	mux.HandleFunc("/atividades/", controllers.GetAtividadePorID)
	mux.HandleFunc("/atividades/cadastrar", controllers.CreateAtividade)
	mux.HandleFunc("/atividades/atualizar", controllers.UpdateAtividade)
	mux.HandleFunc("/atividades/deletar", controllers.DeleteAtividade)
	mux.HandleFunc("/nota/atribuir", controllers.AtribuirNota)
}
