package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

type NotaRepository interface {
	Create(nota *models.Nota) error
	FindByID(id uint) (*models.Nota, error)
	Update(nota *models.Nota) error
	Delete(id uint) error
	FindByAlunoAndAtividade(alunoID, atividadeID uint) (*models.Nota, error)
}

type NotaRepositoryImpl struct{}

func NewNotaRepository() NotaRepository {
	return &NotaRepositoryImpl{}
}

func (r *NotaRepositoryImpl) Create(nota *models.Nota) error {
	return database.DB.Create(nota).Error
}

func (r *NotaRepositoryImpl) FindByID(id uint) (*models.Nota, error) {
	var nota models.Nota
	err := database.DB.First(&nota, id).Error
	if err != nil {
		return nil, err
	}
	return &nota, nil
}

func (r *NotaRepositoryImpl) Update(nota *models.Nota) error {
	return database.DB.Model(nota).Updates(nota).Error
}

func (r *NotaRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Nota{}, id).Error
}

func (r *NotaRepositoryImpl) FindByAlunoAndAtividade(alunoID, atividadeID uint) (*models.Nota, error) {
	var nota models.Nota
	err := database.DB.Where("aluno_id = ? AND atividade_id = ?", alunoID, atividadeID).First(&nota).Error
	if err != nil {
		return nil, err
	}
	return &nota, nil
}
