package models

import (
	"fmt"
	"time"
)

type Artigo struct {
	ID       uint64    `gorm:"primarykey:auto_increment" json:"id"`
	Titulo   string    `gorm:"type:varchar(100)" json:"title"`
	Conteudo string    `json:"content"`
	Criado   time.Time `json:"criado"`
}

func FazerArtigo(a Artigo) {
	bd := fazerCon()
	bd.AutoMigrate(&Artigo{})
	novoArtigo := bd.Create(&a)
	// if err := bd.Create(&a).Error; err != nil {
	// 	log.Fatal("ERRO NO CREATE", err)
	// }
	fmt.Println("artigo feito: ", novoArtigo)
}

func PegarTodosArtigos() []Artigo {
	bd := fazerCon()
	bd.AutoMigrate(&Artigo{})
	var artigos []Artigo
	bd.Find(&artigos)
	// for _, a := range artigos {
	// 	dia := a.Criado.Day()
	// 	fmt.Print(a.ID, "-", a.Titulo, "dia: ", dia)
	// 	fmt.Printf("\n%T %T\n", a.ID, a.Titulo)
	// }
	return artigos
}
