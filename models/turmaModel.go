package models

import (
	"time"
)

type Turma struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Nome      string    `gorm:"not null;size:100"`
	Semestre  int       `gorm:"not null"`
	Ano       int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Turma) TableName() string {
	return "turmas"
}
