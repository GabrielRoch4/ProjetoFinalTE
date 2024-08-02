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
		AllowedOrigins:   []string{"*"},                             // Permite todas as origens
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Métodos permitidos
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true,
	})
	return corsOptions.Handler(handler)
}
