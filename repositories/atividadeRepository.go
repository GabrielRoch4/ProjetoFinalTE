package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

// AtividadeRepository define a interface para as operações de dados de Atividade
type AtividadeRepository interface {
	Create(atividade *models.Atividade) error
	FindByID(id uint) (*models.Atividade, error)
	FindAll() ([]models.Atividade, error)
	Update(atividade *models.Atividade) error
	Delete(id uint) error
}

// AtividadeRepositoryImpl é a implementação do AtividadeRepository
type AtividadeRepositoryImpl struct{}

// NewAtividadeRepository cria uma nova instância do AtividadeRepositoryImpl
func NewAtividadeRepository() AtividadeRepository {
	return &AtividadeRepositoryImpl{}
}

// Create insere uma nova Atividade no banco de dados
func (r *AtividadeRepositoryImpl) Create(atividade *models.Atividade) error {
	return database.DB.Create(atividade).Error
}

// FindByID encontra uma Atividade por ID
func (r *AtividadeRepositoryImpl) FindByID(id uint) (*models.Atividade, error) {
	var atividade models.Atividade
	err := database.DB.Preload("Notas").First(&atividade, id).Error
	if err != nil {
		return nil, err
	}
	return &atividade, nil
}

// FindAll retorna todas as Atividades do banco de dados
func (r *AtividadeRepositoryImpl) FindAll() ([]models.Atividade, error) {
	var atividades []models.Atividade
	err := database.DB.Preload("Notas").Find(&atividades).Error
	return atividades, err
}

// Update atualiza uma Atividade existente no banco de dados
func (r *AtividadeRepositoryImpl) Update(atividade *models.Atividade) error {
	return database.DB.Model(atividade).Updates(atividade).Error
}

// Delete remove uma Atividade do banco de dados por ID
func (r *AtividadeRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Atividade{}, id).Error
}
