package main

import (
	"ProjetoFinal/configs/database"
	"ProjetoFinal/routes"
	"net/http"
)

func main() {

	routes.Router()

	database.DatabaseConnection()

	_ = http.ListenAndServe(":3333", nil)
}
