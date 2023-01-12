package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zennon-sml/JOJORNAL/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var bd *gorm.DB

func FazerCon() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro carregando o arquivo .env")
	}
	//https://www.phpmyadmin.co/ link pro php
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	banco := os.Getenv("DB_NAME")
	tzone := "America/Sao_Paulo"

	//dominio := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usuario, senha, host, porta, banco)
	dominio := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, usuario, senha, banco, porta, tzone)

	bancoDeDados, err := gorm.Open(postgres.Open(dominio), &gorm.Config{})
	if err != nil {
		log.Fatal("não foi possivel estabelecer conexão ", err)
	}
	bd = bancoDeDados
	//MIGRATIONS
	migrations.FazerMigrations(bd)
}

func PegarBD() *gorm.DB {
	return bd
}
