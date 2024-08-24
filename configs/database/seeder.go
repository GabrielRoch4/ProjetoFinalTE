package database

import (
	"ProjetoFinal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Verifique se os dados já existem para evitar duplicação
	var professorCount int64
	db.Model(&models.Professor{}).Count(&professorCount)
	if professorCount > 0 {
		log.Println("Dados já foram inseridos. O seeder não será executado.")
		return
	}

	// Criação de dados de exemplo
	professors := []models.Professor{
		{Nome: "Professor A", Email: "professorA@example.com", CPF: "12345678901"},
		{Nome: "Professor B", Email: "professorB@example.com", CPF: "10987654321"},
	}

	turmas := []models.Turma{
		{Nome: "Turma 1", Semestre: 1, Ano: 2024, ProfessorID: 1},
		{Nome: "Turma 2", Semestre: 2, Ano: 2024, ProfessorID: 2},
	}

	alunos := []models.Aluno{
		{Nome: "Aluno 1", Matricula: 2024001},
		{Nome: "Aluno 2", Matricula: 2024002},
	}

	atividades := []models.Atividade{
		{Nome: "Atividade 1", Valor: 10.00, DataEntrega: time.Now().Add(48 * time.Hour), TurmaID: 1},
		{Nome: "Atividade 2", Valor: 20.00, DataEntrega: time.Now().Add(72 * time.Hour), TurmaID: 2},
	}

	notas := []models.Nota{
		{AlunoID: 1, AtividadeID: 1, Nota: 8.50},
		{AlunoID: 2, AtividadeID: 2, Nota: 7.00},
	}

	// Insere dados
	if err := db.Create(&professors).Error; err != nil {
		log.Fatalf("Erro ao inserir professores: %v", err)
	}
	if err := db.Create(&turmas).Error; err != nil {
		log.Fatalf("Erro ao inserir turmas: %v", err)
	}
	if err := db.Create(&alunos).Error; err != nil {
		log.Fatalf("Erro ao inserir alunos: %v", err)
	}
	if err := db.Create(&atividades).Error; err != nil {
		log.Fatalf("Erro ao inserir atividades: %v", err)
	}
	if err := db.Create(&notas).Error; err != nil {
		log.Fatalf("Erro ao inserir notas: %v", err)
	}

	// Associa alunos a turmas
	var turma1 models.Turma
	var turma2 models.Turma
	var aluno1 models.Aluno
	var aluno2 models.Aluno

	if err := db.First(&turma1, 1).Error; err != nil {
		log.Fatalf("Erro ao buscar turma 1: %v", err)
	}
	if err := db.First(&turma2, 2).Error; err != nil {
		log.Fatalf("Erro ao buscar turma 2: %v", err)
	}
	if err := db.First(&aluno1, 1).Error; err != nil {
		log.Fatalf("Erro ao buscar aluno 1: %v", err)
	}
	if err := db.First(&aluno2, 2).Error; err != nil {
		log.Fatalf("Erro ao buscar aluno 2: %v", err)
	}

	// Adiciona alunos às turmas
	if err := db.Model(&turma1).Association("Alunos").Append(&aluno1, &aluno2); err != nil {
		log.Fatalf("Erro ao associar alunos à turma 1: %v", err)
	}
	if err := db.Model(&turma2).Association("Alunos").Append(&aluno1); err != nil {
		log.Fatalf("Erro ao associar aluno à turma 2: %v", err)
	}
}
