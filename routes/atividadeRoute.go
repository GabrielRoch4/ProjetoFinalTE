package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AtividadeRoutes() {
	http.HandleFunc("/turmas/atividades", controllers.GetAtividade)
	http.HandleFunc("/turmas/atividades/cadastrar", controllers.CreateAtividade)
	http.HandleFunc("/turmas/atividades/atualizar", controllers.UpdateAtividade)
	http.HandleFunc("/turmas/atividades/deletar", controllers.DeleteAtividade)
}
