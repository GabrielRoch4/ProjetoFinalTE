package repositories

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/models"
)

// AlunoRepository define a interface para as operações de dados de Aluno
type AlunoRepository interface {
	Create(aluno *models.Aluno) error
	FindByID(id uint) (*models.Aluno, error)
	FindAll() ([]models.Aluno, error)
	Update(aluno *models.Aluno) error
	Delete(id uint) error
	FindByTurmaID(turmaID uint) ([]models.Aluno, error)
}

// alunoRepositoryImpl é a implementação do AlunoRepository
type alunoRepositoryImpl struct{}

// NewAlunoRepository cria uma nova instância do alunoRepositoryImpl
func NewAlunoRepository() AlunoRepository {
	return &alunoRepositoryImpl{}
}

// Create insere um novo aluno no banco de dados
func (r *alunoRepositoryImpl) Create(aluno *models.Aluno) error {
	return database.DB.Create(aluno).Error
}

// FindByID encontra um aluno por ID, incluindo suas turmas com atividades e notas
func (r *alunoRepositoryImpl) FindByID(id uint) (*models.Aluno, error) {
	var aluno models.Aluno
	err := database.DB.Preload("Turmas").Preload("Turmas.Atividades").Preload("Turmas.Atividades.Notas").First(&aluno, id).Error
	return &aluno, err
}

// FindAll retorna todos os alunos do banco de dados, incluindo suas turmas com atividades e notas
func (r *alunoRepositoryImpl) FindAll() ([]models.Aluno, error) {
	var alunos []models.Aluno
	err := database.DB.Preload("Turmas").Preload("Turmas.Atividades").Preload("Turmas.Atividades.Notas").Find(&alunos).Error
	return alunos, err
}

// Update atualiza um aluno existente no banco de dados
func (r *alunoRepositoryImpl) Update(aluno *models.Aluno) error {
	return database.DB.Model(&models.Aluno{}).Where("id = ?", aluno.ID).Updates(map[string]interface{}{
		"Nome":      aluno.Nome,
		"Matricula": aluno.Matricula,
		"Turmas": aluno.Turmas
	}).Error
}

// Delete remove um aluno do banco de dados por ID
func (r *alunoRepositoryImpl) Delete(id uint) error {
	return database.DB.Delete(&models.Aluno{}, id).Error
}

// FindByTurmaID encontra alunos por ID da turma
func (r *alunoRepositoryImpl) FindByTurmaID(turmaID uint) ([]models.Aluno, error) {
	var alunos []models.Aluno
	err := database.DB.Joins("JOIN turmas_alunos ON alunos.id = turmas_alunos.aluno_id").Where("turmas_alunos.turma_id = ?", turmaID).Find(&alunos).Error
	return alunos, err
}
