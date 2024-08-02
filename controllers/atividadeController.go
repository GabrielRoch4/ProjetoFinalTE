package controllers

import (
	"ProjetoFinal/models"
	"fmt"
	"net/http"
	"time"
)

func GetAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()

	atividade1 := models.Atividade{
		Nome:        "Trabalho Final",
		Valor:       99.50,
		DataEntrega: now,
	}

	fmt.Fprintf(w, "Name: %s, Valor: %.2f, DataEntrega: %s", atividade1.Nome, atividade1.Valor,
		atividade1.DataEntrega.Format("02/01/2006 15:04:05"))
}

func CreateAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Criado")
}

func UpdateAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Atualizado")
}

func DeleteAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Deletado")
}
