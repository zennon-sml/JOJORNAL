package models

import (
	"time"
)

var ModeloDeTempo string = "2006-01-02 15:04:05"

type Artigo struct {
	ID           uint64    `gorm:"primarykey:auto_increment" json:"id"`
	Titulo       string    `gorm:"type:varchar(100)" json:"titulo"`
	Conteudo     string    `json:"conteudo"`
	Criado       time.Time `gorm:"autoCreateTime" json:"criado"`
	UpdatedAt    time.Time
	CriadoModelo string `json: "criadoModelo"`
	AtualizadoModelo string `json: "AtualizadoModelo"`
}

// pra mandar dados como o titulo da pagina
type DadosPagina struct {
	Titulo  string   `json:"titulo"`
	Artigos []Artigo `json:"artigos"`
}
type DadosPagina1 struct {
	Titulo string `json:"titulo"`
	Artigo Artigo `json:"artigos"`
}
