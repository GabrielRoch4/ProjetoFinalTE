package controllers

import (
	"ProjetoFinal/models"
	"fmt"
	"net/http"
)

func GetAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	paulo := models.Aluno{
		Name:      "Paulo",
		Matricula: 1,
	}

	fmt.Fprintf(w, "%+v", paulo)
}

func CreateAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Criado")
}

func UpdateAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Atualizado")
}

func DeleteAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Deletado")
}
