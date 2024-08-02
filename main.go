package main

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/routes"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Configurações do banco de dados
	database.DatabaseConnection()

	// Configuração do roteador com CORS
	router := routes.Router()
	corsHandler := configureCORS(router)

	// Inicia o servidor
	_ = http.ListenAndServe(":3333", corsHandler)
}

// configureCORS configura o middleware CORS
func configureCORS(handler http.Handler) http.Handler {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	return corsOptions.Handler(handler)
}
