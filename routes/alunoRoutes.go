package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func AlunoRoutes() {
	http.HandleFunc("/alunos", controllers.GetAluno)
	http.HandleFunc("/alunos/cadastrar", controllers.CreateAluno)
	http.HandleFunc("/alunos/atualizar", controllers.UpdateAluno)
	http.HandleFunc("/alunos/deletar", controllers.DeleteAluno)
}
