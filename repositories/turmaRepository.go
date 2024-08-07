package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

type TurmaRepository interface {
	Create(turma *models.Turma) error
	FindByID(id uint) (*models.Turma, error)
	FindAll() ([]models.Turma, error)
	Update(turma *models.Turma) error
	Delete(id uint) error
}

type TurmaRepositoryImpl struct{}

func NewTurmaRepository() TurmaRepository {
	return &TurmaRepositoryImpl{}
}

func (r *TurmaRepositoryImpl) Create(turma *models.Turma) error {
	return database.DB.Create(turma).Error
}

func (r *TurmaRepositoryImpl) FindByID(id uint) (*models.Turma, error) {
	var turma models.Turma
	err := database.DB.Preload("Alunos").First(&turma, id).Error
	if err != nil {
		return nil, err
	}
	return &turma, nil
}

func (r *TurmaRepositoryImpl) FindAll() ([]models.Turma, error) {
	var turmas []models.Turma
	err := database.DB.Preload("Alunos").Find(&turmas).Error
	return turmas, err
}

func (r *TurmaRepositoryImpl) Update(turma *models.Turma) error {
	return database.DB.Model(turma).Updates(turma).Error
}

func (r *TurmaRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Turma{}, id).Error
}
