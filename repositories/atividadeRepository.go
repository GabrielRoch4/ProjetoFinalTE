package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

// AtividadeRepository define a interface para as operações de dados de Atividade
type AtividadeRepository interface {
	Create(Atividade *models.Atividade) error
	FindByID(id uint) (*models.Atividade, error)
	FindAll() ([]models.Atividade, error)
	Update(Atividade *models.Atividade) error
	Delete(id uint) error
}

// AtividadeRepositoryImpl é a implementação do AtividadeRepository
type AtividadeRepositoryImpl struct{}

// NewAtividadeRepository cria uma nova instância do AtividadeRepositoryImpl
func NewAtividadeRepository() AtividadeRepository {
	return &AtividadeRepositoryImpl{}
}

// Create insere uma nova Atividade no banco de dados
func (r *AtividadeRepositoryImpl) Create(Atividade *models.Atividade) error {
	return database.DB.Create(Atividade).Error
}

// FindByID encontra uma Atividade por ID
func (r *AtividadeRepositoryImpl) FindByID(id uint) (*models.Atividade, error) {
	var Atividade models.Atividade
	err := database.DB.First(&Atividade, id).Error
	return &Atividade, err
}

// FindAll retorna todas as Atividades do banco de dados
func (r *AtividadeRepositoryImpl) FindAll() ([]models.Atividade, error) {
	var Atividades []models.Atividade
	err := database.DB.Find(&Atividades).Error
	return Atividades, err
}

// Update atualiza uma Atividade existente no banco de dados
func (r *AtividadeRepositoryImpl) Update(Atividade *models.Atividade) error {
	return database.DB.Model(&models.Atividade{}).Where("id = ?", Atividade.ID).Updates(map[string]interface{}{
		"Nome":        Atividade.Nome,
		"Valor":       Atividade.Valor,
		"DataEntrega": Atividade.DataEntrega,
	}).Error
}

// Delete remove uma Atividade do banco de dados por ID
func (r *AtividadeRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Atividade{}, id).Error
}
