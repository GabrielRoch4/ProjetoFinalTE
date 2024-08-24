package handler

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/routes"
	"net/http"

	"github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Configurações do banco de dados
	database.DatabaseConnection()

	// Configuração do roteador com CORS
	router := routes.Router()
	corsHandler := configureCORS(router)

	// Serve o request usando o roteador configurado
	corsHandler.ServeHTTP(w, r)
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
