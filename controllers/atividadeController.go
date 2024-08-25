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

    // Calcula o valor total das atividades da turma
    totalValor, err := atividadeRepo.GetTotalValorByTurmaID(atividade.TurmaID)
    if err != nil {
        http.Error(w, "Erro ao calcular valor total das atividades", http.StatusInternalServerError)
        return
    }

    // Verifica se o valor total das atividades da turma ultrapassa 100 pontos
    if totalValor+atividade.Valor > 100 {
        http.Error(w, "O valor total das atividades da turma não pode ultrapassar 100 pontos", http.StatusBadRequest)
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

	// Calcula o valor total das atividades da turma
	totalValor, err := atividadeRepo.GetTotalValorByTurmaID(existingAtividade.TurmaID)
	if err != nil {
		http.Error(w, "Erro ao calcular valor total das atividades", http.StatusInternalServerError)
		return
	}

	// Subtrai o valor da atividade atual antes de somar o novo valor
	totalValor -= existingAtividade.Valor
	if totalValor+updatedData.Valor > 100 {
		http.Error(w, "O valor total das atividades da turma não pode ultrapassar 100 pontos", http.StatusBadRequest)
		return
	}

	// Atualiza os campos necessários
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
        AlunoID     uint    `json:"AlunoID"`
        AtividadeID uint    `json:"AtividadeID"`
        Nota        float64 `json:"Nota"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

	atividade, err := atividadeRepo.FindByID(input.AtividadeID)
    if err != nil {
        http.Error(w, "Erro ao buscar atividade", http.StatusInternalServerError)
        return
    }

    if input.Nota > atividade.Valor {
        http.Error(w, "Nota não pode ser maior que o valor máximo da atividade", http.StatusBadRequest)
        return
    }

    nota, err := notaRepo.FindByAlunoAndAtividade(input.AlunoID, input.AtividadeID)
    if err != nil {
        http.Error(w, "Erro ao buscar nota", http.StatusInternalServerError)
        return
    }

    if nota != nil {
        nota.Nota = input.Nota
        if err := notaRepo.Update(nota); err != nil {
            http.Error(w, "Erro ao atualizar nota", http.StatusInternalServerError)
            return
        }
    } else {
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

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "Nota atribuída com sucesso"})
}