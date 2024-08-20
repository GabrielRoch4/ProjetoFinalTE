package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

var alunoRepo = repositories.NewAlunoRepository()

// GetAluno - Obter todos os alunos, incluindo turmas com atividades e notas
func GetAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	alunos, err := alunoRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter alunos", http.StatusInternalServerError)
		return
	}

	if len(alunos) == 0 {
		http.Error(w, "Não há alunos cadastrados", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alunos)
}

// GetAlunoPorID - Obter um aluno por ID, incluindo turmas com atividades e notas
func GetAlunoPorID(w http.ResponseWriter, r *http.Request) {
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

	aluno, err := alunoRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	if aluno == nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aluno)
}

// CreateAluno - Criar um novo aluno
func CreateAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var aluno models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&aluno); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := alunoRepo.Create(&aluno); err != nil {
		http.Error(w, "Erro ao criar aluno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aluno)
}

// UpdateAluno - Atualizar um aluno existente
func UpdateAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var updatedData models.Aluno
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	existingAluno, err := alunoRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar aluno", http.StatusInternalServerError)
		return
	}

	if existingAluno == nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	existingAluno.Nome = updatedData.Nome
	existingAluno.Matricula = updatedData.Matricula

	if err := alunoRepo.Update(existingAluno); err != nil {
		http.Error(w, "Erro ao atualizar aluno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Atualizado com sucesso!")
}

// DeleteAluno - Excluir um aluno por ID
func DeleteAluno(w http.ResponseWriter, r *http.Request) {
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

	aluno, err := alunoRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar aluno", http.StatusInternalServerError)
		return
	}

	if aluno == nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	err = alunoRepo.Delete(uint(id))
	if err != nil {
		http.Error(w, "Erro ao deletar aluno", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}
