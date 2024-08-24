package models

import (
	"time"
)

type Turma struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Nome        string `gorm:"not null;size:100"`
	Semestre    int    `gorm:"not null"`
	Ano         int    `gorm:"not null"`
	ProfessorID uint   `gorm:"not null"`
	Atividades  []Atividade
	Alunos      []Aluno   `gorm:"many2many:turmas_alunos;"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (Turma) TableName() string {
	return "turmas"
}
