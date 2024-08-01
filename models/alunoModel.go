package models

import (
	"time"
)

type Aluno struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"not null;size:100"`
	Matricula int        `gorm:"not null;unique;size:100"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (Aluno) TableName() string {
	return "alunos"
}
