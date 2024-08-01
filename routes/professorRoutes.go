package routes

import (
	"ProjetoFinal/controllers"
	"net/http"
)

func ProfessorRoutes() {
	http.HandleFunc("/professores", controllers.GetProfessor)
	http.HandleFunc("/professores/", controllers.GetProfessorPorID)
	http.HandleFunc("/professores/cadastrar", controllers.CreateProfessor)
	http.HandleFunc("/professores/atualizar", controllers.UpdateProfessor)
	http.HandleFunc("/professores/deletar", controllers.DeleteProfessor)
}
