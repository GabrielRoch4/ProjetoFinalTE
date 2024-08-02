package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AlunoRoutes(mux *http.ServeMux) {
	http.HandleFunc("/alunos", controllers.GetAluno)
	http.HandleFunc("/alunos/", controllers.GetAlunoPorID)
	http.HandleFunc("/alunos/cadastrar", controllers.CreateAluno)
	http.HandleFunc("/alunos/atualizar", controllers.UpdateAluno)
	http.HandleFunc("/alunos/deletar", controllers.DeleteAluno)
}
