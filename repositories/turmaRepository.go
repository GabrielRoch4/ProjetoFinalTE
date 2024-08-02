package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

type TurmaRepository interface {
	Create(Turma *models.Turma) error
	FindByID(id uint) (*models.Turma, error)
	FindAll() ([]models.Turma, error)
	Update(Turma *models.Turma) error
	Delete(id uint) error
}

// TurmaRepositoryImpl é a implementação do TurmaRepository
type TurmaRepositoryImpl struct{}

// NewTurmaRepository cria uma nova instância do TurmaRepositoryImpl
func NewTurmaRepository() TurmaRepository {
	return &TurmaRepositoryImpl{}
}

// Create insere uma nova Turma no banco de dados
func (r *TurmaRepositoryImpl) Create(Turma *models.Turma) error {
	return database.DB.Create(Turma).Error
}

// FindByID encontra uma Turma por ID
func (r *TurmaRepositoryImpl) FindByID(id uint) (*models.Turma, error) {
	var Turma models.Turma
	err := database.DB.First(&Turma, id).Error
	return &Turma, err
}

// FindAll retorna todas as Turmas do banco de dados
func (r *TurmaRepositoryImpl) FindAll() ([]models.Turma, error) {
	var Turmas []models.Turma
	err := database.DB.Find(&Turmas).Error
	return Turmas, err
}

// Update atualiza uma Turma existente no banco de dados
func (r *TurmaRepositoryImpl) Update(Turma *models.Turma) error {
	return database.DB.Model(&models.Turma{}).Where("id = ?", Turma.ID).Updates(map[string]interface{}{
		"Nome":     Turma.Nome,
		"Semestre": Turma.Semestre,
		"Ano":      Turma.Ano,
	}).Error
}

// Delete remove uma Turma do banco de dados por ID
func (r *TurmaRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Turma{}, id).Error
}
