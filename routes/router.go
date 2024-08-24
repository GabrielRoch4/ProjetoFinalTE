package routes

import (
	"net/http"
)

// Router configura todas as rotas da aplicação
func Router() http.Handler {
	mux := http.NewServeMux() // Cria um novo roteador
	ProfessorRoutes(mux)
	TurmaRoutes(mux)
	AlunoRoutes(mux)
	AtividadeRoutes(mux)
	return mux
}
