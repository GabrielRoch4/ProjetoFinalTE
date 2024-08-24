package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

var turmaRepo = repositories.NewTurmaRepository()

func GetTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	turmas, err := turmaRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter turmas", http.StatusInternalServerError)
		return
	}

	if len(turmas) == 0 {
		http.Error(w, "Não há turmas cadastradas", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(turmas)
}

func GetTurmaPorID(w http.ResponseWriter, r *http.Request) {
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

	turma, err := turmaRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar turma", http.StatusInternalServerError)
		return
	}

	if turma == nil {
		http.Error(w, "Turma não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(turma)
}

func CreateTurma(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var turma models.Turma
	if err := json.NewDecoder(r.Body).Decode(&turma); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := turmaRepo.Create(&turma); err != nil {
		http.Error(w, "Erro ao criar turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turma)
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

	existingTurma, err := turmaRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar turma", http.StatusInternalServerError)
		return
	}

	if existingTurma == nil {
		http.Error(w, "Turma não encontrada", http.StatusNotFound)
		return
	}

	existingTurma.Nome = updatedData.Nome
	existingTurma.Semestre = updatedData.Semestre
	existingTurma.Ano = updatedData.Ano

	if err := turmaRepo.Update(existingTurma); err != nil {
		http.Error(w, "Erro ao atualizar turma", http.StatusInternalServerError)
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

	turma, err := turmaRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar turma", http.StatusInternalServerError)
		return
	}

	if turma == nil {
		http.Error(w, "Turma não encontrada", http.StatusNotFound)
		return
	}

	if err := turmaRepo.Delete(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}

func AdicionarAluno(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		TurmaID uint `json:"turma_id"`
		AlunoID uint `json:"aluno_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	turma, err := turmaRepo.FindByID(input.TurmaID)
	if err != nil {
		http.Error(w, "Erro ao buscar turma", http.StatusInternalServerError)
		return
	}

	if turma == nil {
		http.Error(w, "Turma não encontrada", http.StatusNotFound)
		return
	}

	aluno, err := alunoRepo.FindByID(input.AlunoID) // Usando o repositório de alunos
	if err != nil {
		http.Error(w, "Erro ao buscar aluno", http.StatusInternalServerError)
		return
	}

	if aluno == nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	turma.Alunos = append(turma.Alunos, *aluno)

	if err := turmaRepo.Update(turma); err != nil {
		http.Error(w, "Erro ao adicionar aluno à turma", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Aluno adicionado à turma com sucesso!")
}
