package controllers

import (
	"ProjetoFinal/models"
	"ProjetoFinal/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

var atividadeRepo = repositories.NewAtividadeRepository()
var notaRepo = repositories.NewNotaRepository()

func GetAtividades(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	atividades, err := atividadeRepo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao obter atividades", http.StatusInternalServerError)
		return
	}

	if len(atividades) == 0 {
		http.Error(w, "Não há atividades cadastradas", http.StatusNotFound)
		return
	}

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

	atividade, err := atividadeRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar atividade", http.StatusInternalServerError)
		return
	}

	if atividade == nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

	if err := atividadeRepo.Create(&atividade); err != nil {
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

	existingAtividade, err := atividadeRepo.FindByID(updatedData.ID)
	if err != nil {
		http.Error(w, "Erro ao buscar atividade", http.StatusInternalServerError)
		return
	}

	if existingAtividade == nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	existingAtividade.Nome = updatedData.Nome
	existingAtividade.Valor = updatedData.Valor
	existingAtividade.DataEntrega = updatedData.DataEntrega

	if err := atividadeRepo.Update(existingAtividade); err != nil {
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

	atividade, err := atividadeRepo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Erro ao buscar atividade", http.StatusInternalServerError)
		return
	}

	if atividade == nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	if err := atividadeRepo.Delete(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar atividade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deletado com sucesso!")
}

func AtribuirNota(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		AlunoID     uint    `json:"aluno_id"`
		AtividadeID uint    `json:"atividade_id"`
		Nota        float64 `json:"nota"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Validação dos dados
	if input.Nota < 0 {
		http.Error(w, "Valor da nota inválido", http.StatusBadRequest)
		return
	}

	// Verificar se a atividade existe
	atividade, err := atividadeRepo.FindByID(input.AtividadeID)
	if err != nil {
		http.Error(w, "Erro ao buscar atividade", http.StatusInternalServerError)
		return
	}

	if atividade == nil {
		http.Error(w, "Atividade não encontrada", http.StatusNotFound)
		return
	}

	// Verificar se o aluno existe
	aluno, err := alunoRepo.FindByID(input.AlunoID)
	if err != nil {
		http.Error(w, "Erro ao buscar aluno", http.StatusInternalServerError)
		return
	}

	if aluno == nil {
		http.Error(w, "Aluno não encontrado", http.StatusNotFound)
		return
	}

	// Verificar se a nota já existe para o aluno e a atividade
	nota, err := notaRepo.FindByAlunoAndAtividade(input.AlunoID, input.AtividadeID)
	if err != nil && err.Error() != "record not found" {
		http.Error(w, "Erro ao buscar nota", http.StatusInternalServerError)
		return
	}

	if nota != nil {
		// Atualizar nota existente
		nota.Nota = input.Nota
		if err := notaRepo.Update(nota); err != nil {
			http.Error(w, "Erro ao atualizar nota", http.StatusInternalServerError)
			return
		}
	} else {
		// Criar nova nota
		nota = &models.Nota{
			AlunoID:     input.AlunoID,
			AtividadeID: input.AtividadeID,
			Nota:        input.Nota,
		}
		if err := notaRepo.Create(nota); err != nil {
			http.Error(w, "Erro ao atribuir nota", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Nota atribuída com sucesso!")
}
