package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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

func FormArtigo(c *gin.Context) {
	dados := DadosPagina{Titulo: "fazer artigo"}
	c.HTML(200, "fazerArtigo.html", dados)
}

func FormAtualizarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := fazerCon()
	var artigo Artigo
	bd.First(&artigo, id)
	dados := DadosPagina1{Titulo: "Atualizar artigo", Artigo: artigo}
	c.HTML(200, "atualizarArtigo.html", dados)
}

func FazerArtigo(c *gin.Context) {
	var artigo Artigo
	titulo := c.DefaultPostForm("titulo", "not found")
	conteudo := c.DefaultPostForm("conteudo", "not found")
	artigo.Titulo = titulo
	artigo.Conteudo = conteudo
	bd := fazerCon()
	bd.Debug().Create(&artigo)

	c.HTML(200, "artigo.html", artigo)
}

func PegarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := fazerCon()
	var artigo Artigo
	bd.Debug().First(&artigo, id)
	artigo.CriadoModelo = artigo.Criado.Format(ModeloDeTempo)
	artigo.AtualizadoModelo = artigo.UpdatedAt.Format(ModeloDeTempo)
	c.HTML(200, "artigo.html", artigo)
}

func PegarTodosArtigos(c *gin.Context) {
	bd := fazerCon()
	var artigos []Artigo
	bd.Debug().Order("id desc").Find(&artigos)
	for i, _ := range artigos {
		artigos[i].CriadoModelo = artigos[i].Criado.Format(ModeloDeTempo)
	}
	fmt.Println(artigos[1].CriadoModelo)
	pagina := DadosPagina{Titulo: "todos os artigo", Artigos: artigos}
	// c.IndentedJSON(200, artigos)
	c.HTML(200, "jojornal.html", pagina)
}

//TODO em vez de pegar o id pela url mandar todo a struct
func ApagarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := fazerCon()
	var artigo Artigo
	bd.Debug().Delete(&artigo, id)

	PegarTodosArtigos(c)
}

func AtualizarArtigo(c *gin.Context) {

	id := c.Param("id")
	titulo := c.DefaultPostForm("titulo", "not found")
	conteudo := c.DefaultPostForm("conteudo", "not found")
	var artigo Artigo
	bd := fazerCon()
	bd.First(&artigo, id)
	artigo.Titulo = titulo
	artigo.Conteudo = conteudo
	bd.Debug().Save(artigo)

	fmt.Println(id, titulo, conteudo, artigo)
	c.HTML(200, "artigo.html", artigo)
}
