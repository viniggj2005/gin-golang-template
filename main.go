package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/viniggj2005/api-rest-go/routes"
)

// @title           API de Template
// @version         1.0
// @description     API REST em Go Template
// @host            localhost:3000
// @BasePath        /
// @schemes         http
// @tag.name        Template
// @tag.description Operações relacionadas a Template
func main() {
	routes.InitValidator()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar env .env file: %s", err)
	} //configuração para carregar as variáveis de ambiente
	// database.ConectaDB()
	routes.HanddleRequests()

}
