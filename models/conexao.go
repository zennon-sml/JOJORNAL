package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var artigo Artigo

func fazerCon() *gorm.DB { //*sql.DB
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	banco := os.Getenv("DB_NAME")

	dominio := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usuario, senha, host, porta, banco)

	bd, err := gorm.Open(mysql.Open(dominio), &gorm.Config{})
	if err != nil {
		log.Fatal("não foi possivel estabelecer conexão ", err)
	}
	bd.AutoMigrate(&artigo)
	return bd
}
