package models

import (
	"time"
)

type Professor struct {
	ID        uint       `gorm:"primaryKey;autoincrement"`
	Nome      string     `gorm:"not null;size:50"`
	Email     string     `gorm:"not null;unique;size:100"`
	CPF       string     `gorm:"not null;unique;size:14"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (Professor) TableName() string {
	return "professores"
}
