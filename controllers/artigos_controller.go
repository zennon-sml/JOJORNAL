package controllers

import(
    "github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/database"
	"github.com/zennon-sml/JOJORNAL/models"
)

func FormArtigo(c *gin.Context) {
	dados := models.DadosPagina{Titulo: "fazer artigo"}
	c.HTML(200, "fazerArtigo.html", dados)
}

func FormAtualizarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := database.PegarBD()
	var artigo models.Artigo
	bd.First(&artigo, id)
	dados := models.DadosPagina1{Titulo: "Atualizar artigo", Artigo: artigo}
	c.HTML(200, "atualizarArtigo.html", dados)
}

func FazerArtigo(c *gin.Context) {
	var artigo models.Artigo
	titulo := c.DefaultPostForm("titulo", "not found")
	conteudo := c.DefaultPostForm("conteudo", "not found")
	artigo.Titulo = titulo
	artigo.Conteudo = conteudo
	bd := database.PegarBD()
	bd.Debug().Create(&artigo)
	FormatarTempo(&artigo)
	c.HTML(200, "artigo.html", artigo)
}

func PegarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := database.PegarBD()
	var artigo models.Artigo
	bd.Debug().First(&artigo, id)
	FormatarTempo(&artigo)

    c.HTML(200, "artigo.html", artigo)
}

func PegarTodosArtigos(c *gin.Context) {
	bd := database.PegarBD()
	var artigos []models.Artigo
	bd.Debug().Order("id desc").Find(&artigos)
	for i, _ := range artigos {
	    FormatarTempo(&artigos[i])
    }
	pagina := models.DadosPagina{Titulo: "todos os artigo", Artigos: artigos}
	// c.IndentedJSON(200, artigos)
	c.HTML(200, "jojornal.html", pagina)
}

//TODO em vez de pegar o id pela url mandar todo a struct
func ApagarArtigo(c *gin.Context) {
	id := c.Param("id")
	bd := database.PegarBD()
	var artigo models.Artigo
	bd.Debug().Delete(&artigo, id)

	PegarTodosArtigos(c)
}

func AtualizarArtigo(c *gin.Context) {
	id := c.Param("id")
	titulo := c.DefaultPostForm("titulo", "not found")
	conteudo := c.DefaultPostForm("conteudo", "not found")
	var artigo models.Artigo
	bd := database.PegarBD()
	bd.First(&artigo, id)
	artigo.Titulo = titulo
	artigo.Conteudo = conteudo
	bd.Debug().Save(artigo)
    FormatarTempo(&artigo) 


	c.HTML(200, "artigo.html", artigo)
}

func FormatarTempo(a *models.Artigo){
	a.CriadoModelo = a.Criado.Format(models.ModeloDeTempo)
	a.AtualizadoModelo = a.UpdatedAt.Format(models.ModeloDeTempo)
}
