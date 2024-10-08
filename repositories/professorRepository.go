package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

// ProfessorRepository define a interface para as operações de dados de Professor
type ProfessorRepository interface {
	Create(professor *models.Professor) error
	FindByID(id uint) (*models.Professor, error)
	FindAll() ([]models.Professor, error)
	Update(professor *models.Professor) error
	Delete(id uint) error
}

// professorRepositoryImpl é a implementação do ProfessorRepository
type professorRepositoryImpl struct{}

// NewProfessorRepository cria uma nova instância do professorRepositoryImpl
func NewProfessorRepository() ProfessorRepository {
	return &professorRepositoryImpl{}
}

// Create insere um novo professor no banco de dados
func (r *professorRepositoryImpl) Create(professor *models.Professor) error {
	return database.DB.Create(professor).Error
}

// FindByID encontra um professor por ID
func (r *professorRepositoryImpl) FindByID(id uint) (*models.Professor, error) {
	var professor models.Professor

	err := database.DB.
		Preload("Turmas").
		Preload("Turmas.Atividades").
		Preload("Turmas.Alunos").
		First(&professor, id).Error

	if err != nil {
		return nil, err
	}
	return &professor, nil
}

// FindAll retorna todos os professores do banco de dados
func (r *professorRepositoryImpl) FindAll() ([]models.Professor, error) {
	var professors []models.Professor

	err := database.DB.
		Preload("Turmas").
		Preload("Turmas.Atividades").
		Preload("Turmas.Alunos").
		Find(&professors).Error

	return professors, err
}

// Update atualiza um professor existente no banco de dados
func (r *professorRepositoryImpl) Update(professor *models.Professor) error {
	return database.DB.Model(professor).Updates(professor).Error
}

// Delete remove um professor do banco de dados por ID
func (r *professorRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Professor{}, id).Error
}
