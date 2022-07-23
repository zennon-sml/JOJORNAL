package models

import (
	"time"
)

var ModeloDeTempo string = "2006-01-02 15:04:05"

type Administrador struct{
	ID uint32    `gorm:"primarykey:auto_increment" json:"id"`
	Nome string  `gorm:"type:varchar(100)" json:"nome"`
	Email string `gorm:"type:varchar(200)" json:"email"`
	Senha string `gorm:"type:varchar(50) json:"-"`
	Entrada time.Time `gorm:"autoCreateTime" json:"entrada"`
	Artigos []Artigo `json:"artigos"` //define a foreign key para os artigos

}

type Artigo struct {
	ID           uint32`gorm:"primarykey:auto_increment" json:"id"`
	Titulo       string    `gorm:"type:varchar(100)" json:"titulo"`
	Conteudo     string    `json:"conteudo"`
	ImagemUrl	 string    `gorm:"type:varchar(255) json:"imagem_url"`
	AdministradorID     uint32 `json:"admin_id`
	Criado       time.Time `gorm:"autoCreateTime" json:"criado"`
	UpdatedAt    time.Time
	CriadoModelo string    `json: "criadoModelo"`
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
