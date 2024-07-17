package models

import "time"

type Atividade struct {
	Name        string
	Valor       float64
	DataEntrega time.Time
}
