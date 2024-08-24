package models

import (
	"time"
)

type Aluno struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Nome      string     `gorm:"not null;size:100"`
	Matricula int        `gorm:"not null;unique;size:100"`
	Turmas    []Turma    `gorm:"many2many:turmas_alunos;"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (Aluno) TableName() string {
	return "alunos"
}
