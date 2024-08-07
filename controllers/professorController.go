package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

// Repositório de professor
var professorRepo = repositories.NewProfessorRepository()

// GetProfessores retorna todos os professores cadastrados
func GetProfessores(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	professores, err := professorRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter professores", http.StatusInternalServerError)
		return
	}

	if len(professores) == 0 {
		http.Error(w, "Não há professores cadastrados", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(professores)
}

// GetProfessorPorID retorna um professor pelo seu ID
func GetProfessorPorID(w http.ResponseWriter, r *http.Request) {
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

	professor, err := professorRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar professor", http.StatusInternalServerError)
		return
	}

	if professor == nil {
		http.Error(w, "Professor não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(professor)
}

// CreateProfessor cria um novo professor
func CreateProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var professor models.Professor
	if err := json.NewDecoder(r.Body).Decode(&professor); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := professorRepo.Create(&professor); err != nil {
		http.Error(w, "Erro ao criar professor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(professor)
}

// UpdateProfessor atualiza um professor existente
func UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var updatedData models.Professor
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	existingProfessor, err := professorRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar professor", http.StatusInternalServerError)
		return
	}

	if existingProfessor == nil {
		http.Error(w, "Professor não encontrado", http.StatusNotFound)
		return
	}

	existingProfessor.Nome = updatedData.Nome
	existingProfessor.Email = updatedData.Email
	existingProfessor.CPF = updatedData.CPF

	if err := professorRepo.Update(existingProfessor); err != nil {
		http.Error(w, "Erro ao atualizar professor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Atualizado com sucesso!")
}

// DeleteProfessor remove um professor pelo ID
func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
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

	professor, err := professorRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar professor", http.StatusInternalServerError)
		return
	}

	if professor == nil {
		http.Error(w, "Professor não encontrado", http.StatusNotFound)
		return
	}

	if err := professorRepo.Delete(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar professor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}
