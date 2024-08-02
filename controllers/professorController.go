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

func GetProfessor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Exemplo: obter todos os professores
	professors, err := professorRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter professores", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professors)
}

func GetProfessorPorID(w http.ResponseWriter, r *http.Request) {
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

	professor, err := professorRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar professor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professor)
}

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

	// Buscando o professor atual no banco de dados
	existingProfessor, err := professorRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar professor", http.StatusInternalServerError)
		return
	}

	// Atualizando apenas os campos desejados
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

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
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
	professor, err := professorRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Professor não encontrado", http.StatusNotFound)
		return
	}
	if professor == nil {
		http.Error(w, "Professor não encontrado", http.StatusNotFound)
		return
	}

	// Tentativa de exclusão do professor pelo ID
	err = professorRepo.Delete(uint(id))
	if err != nil {
		http.Error(w, "Erro ao deletar professor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}
