package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AlunoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/alunos", controllers.GetAluno)
	mux.HandleFunc("/alunos/", controllers.GetAlunoPorID)
	mux.HandleFunc("/alunos/cadastrar", controllers.CreateAluno)
	mux.HandleFunc("/alunos/atualizar", controllers.UpdateAluno)
	mux.HandleFunc("/alunos/deletar", controllers.DeleteAluno)
}
