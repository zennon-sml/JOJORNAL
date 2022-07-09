package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Artigo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func pegarTodos() {
	//bd := fazerCon()

}

func FazerArtigo(a Artigo) {
	bd := fazerCon()
	bd.AutoMigrate(&Artigo{})
	bd.Create(&a)
	if err := bd.Create(&a).Error; err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("artigo feito")
	}
}
