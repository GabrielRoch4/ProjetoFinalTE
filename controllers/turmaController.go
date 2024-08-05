package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

var TurmaRepo = repositories.NewTurmaRepository()

func GetTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	Turmas, err := TurmaRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter Turmas", http.StatusInternalServerError)
		return
	}

	if len(Turmas) == 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Não há turmas cadastradas", http.StatusNotFound)
		return
	}

	// Se houver alunos, retornará a lista normalmente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Turmas)
}

func GetTurmaPorID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtendo o ID a partir dos parâmetros de consulta
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

	Turma, err := TurmaRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Turma não encontrada!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Turma)
}

func CreateTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var Turma models.Turma
	if err := json.NewDecoder(r.Body).Decode(&Turma); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := TurmaRepo.Create(&Turma); err != nil {
		http.Error(w, "Erro ao criar Turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Turma)
}

func UpdateTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var updatedData models.Turma
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	existingTurma, err := TurmaRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Turma não encontrado", http.StatusNotFound)
		return
	}

	existingTurma.Nome = updatedData.Nome
	existingTurma.Semestre = updatedData.Semestre
	existingTurma.Ano = updatedData.Ano

	if err := TurmaRepo.Update(existingTurma); err != nil {
		http.Error(w, "Erro ao atualizar Turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Atualizado com sucesso!")
}

func DeleteTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtendo o ID a partir dos parâmetros de consulta
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

	// Verificando se o professor existe
	Turma, err := TurmaRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Turma não encontrado", http.StatusNotFound)
		return
	}
	if Turma == nil {
		http.Error(w, "Turma não encontrado", http.StatusNotFound)
		return
	}

	// Tentativa de exclusão do professor pelo ID
	err = TurmaRepo.Delete(uint(id))
	if err != nil {
		http.Error(w, "Erro ao deletar Turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}
