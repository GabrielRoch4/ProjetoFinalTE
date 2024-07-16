package handlers


import (
    "ProjetoFinal/controllers"
    "net/http"
)

func ProfessorHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        controllers.GetProfessor(w, r)
    case http.MethodPost:
        controllers.CreateProfessor(w, r)
    case http.MethodPut:
        controllers.UpdateProfessor(w, r)
    case http.MethodDelete:
        controllers.DeleteProfessor(w, r)
    default:
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
    }
}