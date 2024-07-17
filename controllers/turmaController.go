package controllers

import (
	"ProjetoFinal/models"
	"fmt"
	"net/http"
)

func GetTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	ads := models.Turma{
		Name:     "ADS",
		Semestre: 1,
		Ano:      2024,
		Professor: models.Professor{
			Name:  "André",
			Email: "andre@gmail.com",
			CPF:   "12345678910",
		},
	}

	fmt.Fprintf(w, "%+v", ads)
}

func CreateTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Criado")
}

func UpdateTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Atualizado")
}

func DeleteTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Deletado")
}
