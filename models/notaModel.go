package models

import (
	"time"
)

type Nota struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	AlunoID     uint       `gorm:"not null"`
	AtividadeID uint       `gorm:"not null"`
	Nota        float64    `gorm:"type:decimal(5,2);not null"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
}

func (Nota) TableName() string {
	return "notas"
}
