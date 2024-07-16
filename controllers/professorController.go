package controllers

import (
	"ProjetoFinal/models"
	"fmt"
	"net/http"
)

func GetProfessor(w http.ResponseWriter, r *http.Request) {
	gabriel := models.Professor{
		Name:  "Gabriel",
		Email: "gabriel@gmail.com",
		CPF:   "12312312311",
	}

	fmt.Fprintf(w, "%+v", gabriel)
}

func CreateProfessor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Criado")
}

func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Atualizado")
}

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deletado")
}
