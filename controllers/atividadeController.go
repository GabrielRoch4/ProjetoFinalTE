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

    var inputs []struct {
        AlunoID     uint    `json:"AlunoID"`
        AtividadeID uint    `json:"AtividadeID"`
        Nota        float64 `json:"Nota"`
    }

    if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    var erros []string

    for _, input := range inputs {
        if input.Nota < 0 {
            erros = append(erros, "Valor da nota inválido para AlunoID: "+strconv.Itoa(int(input.AlunoID))+" e AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
            continue
        }

        atividade, err := atividadeRepo.FindByID(input.AtividadeID)
        if err != nil {
            erros = append(erros, "Erro ao buscar atividade para AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
            continue
        }

        if atividade == nil {
            erros = append(erros, "Atividade não encontrada para AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
            continue
        }

        if input.Nota > atividade.Valor {
            erros = append(erros, "Nota não pode ser maior que o valor da atividade para AlunoID: "+strconv.Itoa(int(input.AlunoID))+" e AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
            continue
        }

        aluno, err := alunoRepo.FindByID(input.AlunoID)
        if err != nil {
            erros = append(erros, "Erro ao buscar aluno para AlunoID: "+strconv.Itoa(int(input.AlunoID)))
            continue
        }

        if aluno == nil {
            erros = append(erros, "Aluno não encontrado para AlunoID: "+strconv.Itoa(int(input.AlunoID)))
            continue
        }

        nota, err := notaRepo.FindByAlunoAndAtividade(input.AlunoID, input.AtividadeID)
        if err != nil && err.Error() != "record not found" {
            erros = append(erros, "Erro ao buscar nota para AlunoID: "+strconv.Itoa(int(input.AlunoID))+" e AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
            continue
        }

        if nota != nil {
            nota.Nota = input.Nota
            if err := notaRepo.Update(nota); err != nil {
                erros = append(erros, "Erro ao atualizar nota para AlunoID: "+strconv.Itoa(int(input.AlunoID))+" e AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
                continue
            }
        } else {
            nota = &models.Nota{
                AlunoID:     input.AlunoID,
                AtividadeID: input.AtividadeID,
                Nota:        input.Nota,
            }
            if err := notaRepo.Create(nota); err != nil {
                erros = append(erros, "Erro ao atribuir nota para AlunoID: "+strconv.Itoa(int(input.AlunoID))+" e AtividadeID: "+strconv.Itoa(int(input.AtividadeID)))
                continue
            }
        }
    }

    if len(erros) > 0 {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(erros)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Notas atribuídas com sucesso!")
}