package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Artigo struct {
	ID       uint64    `gorm:"primarykey:auto_increment" json:"id"`
	Titulo   string    `gorm:"type:varchar(100)" json:"titulo"`
	Conteudo string    `json:"conteudo"`
	Criado   time.Time `gorm:"autoCreateTime" json:"criado"`
}

// pra mandar dados como o titulo da pagina
type DadosPagina struct {
	Titulo string `json:"titulo"`
}

func FazerArtigo(c *gin.Context) {
	bd := fazerCon()

	var artigo Artigo

	if err := c.BindJSON(&artigo); err != nil {
		log.Fatal("ERRO FAZENDO ARTIGO, ", err)
	}
	bd.Create(&artigo)
	c.IndentedJSON(201, artigo)
}

func PegarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := fazerCon()
	var artigo Artigo
	bd.First(&artigo, id)

	c.IndentedJSON(200, artigo)
}

func PegarTodosArtigos(c *gin.Context) {
	bd := fazerCon()
	var artigos []Artigo
	bd.Find(&artigos)
	c.IndentedJSON(200, artigos)
}

//TODO em vez de pegar o id pela url mandar todo a struct
func ApagarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := fazerCon()
	var artigo Artigo
	bd.Delete(&artigo, id)

	c.IndentedJSON(204, artigo)
}

func AtualizarArtigo(c *gin.Context) {
	bd := fazerCon()
	var artigoNovo Artigo
	if err := c.BindJSON(&artigoNovo); err != nil {
		log.Fatal("ERRO FAZENDO ARTIGO, ", err)
	}
	var artigoVelho Artigo
	bd.First(&artigoVelho, artigoNovo.ID)
	artigoVelho.Titulo = artigoNovo.Titulo
	artigoVelho.Conteudo = artigoNovo.Conteudo
	bd.Save(&artigoVelho)

	c.IndentedJSON(201, artigoVelho)
}

func HTMLExemplo(c *gin.Context) {
	titulo := DadosPagina{
		Titulo: "123 testando",
	} //json que vai ser mandado pra pagina
	c.HTML(200, "artigo.html", titulo)
	//   status   qual HTML     JSON
}
