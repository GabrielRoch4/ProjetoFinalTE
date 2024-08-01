package models

import (
	"time"
)

type Atividade struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Name        string     `gorm:"not null;size:100"`
	Valor       float64    `gorm:"type:decimal(5,2);not null"`
	DataEntrega time.Time  `gorm:"not null"`
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
}

func (Atividade) TableName() string {
	return "atividades"
}
