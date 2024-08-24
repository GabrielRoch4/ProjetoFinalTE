// database/database.go
package database

import (
	"ProjetoFinal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/projeto_final_te?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.Professor{},
		&models.Aluno{},
		&models.Atividade{},
		&models.Turma{},
		&models.Nota{},
	)

	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
}
