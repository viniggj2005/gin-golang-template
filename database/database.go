package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	connection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port,
	)
	// connection := "host=pg-go-rest user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Erro ao conectar com o db")
	}
	DB.AutoMigrate() //aqui vai criar as tabelas do banco de dados ex:&models.User{}
}
