package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

var AtividadeRepo = repositories.NewAtividadeRepository()

func GetAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	atividades, err := AtividadeRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter atividades", http.StatusInternalServerError)
		return
	}

	if len(atividades) == 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Não há atividades cadastradas", http.StatusNotFound)
		return
	}

	// Se houver alunos, retornará a lista normalmente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(atividades)
}

func GetAtividadePorID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	atividade, err := AtividadeRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atividade)
}

func CreateAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var atividade models.Atividade
	if err := json.NewDecoder(r.Body).Decode(&atividade); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := AtividadeRepo.Create(&atividade); err != nil {
		http.Error(w, "Erro ao criar atividade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(atividade)
}

func UpdateAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var updatedData models.Atividade
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	existingAtividade, err := AtividadeRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	existingAtividade.Nome = updatedData.Nome
	existingAtividade.Valor = updatedData.Valor
	existingAtividade.DataEntrega = updatedData.DataEntrega

	if err := AtividadeRepo.Update(existingAtividade); err != nil {
		http.Error(w, "Erro ao atualizar atividade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Atualizado com sucesso!")
}

func DeleteAtividade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	atividade, err := AtividadeRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	if atividade == nil {
		http.Error(w, "Turma não encontrado", http.StatusNotFound)
		return
	}

	if err := AtividadeRepo.Delete(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar atividade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}
