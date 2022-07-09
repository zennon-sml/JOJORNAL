package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

var artigos []models.Artigo

func getArtigos(c *gin.Context) {
	c.JSON(http.StatusOK, artigos)
}

func main() {
	var art models.Artigo
	var titulo, conteudo string
	fmt.Print("Titulo: ")
	fmt.Scanf("%s", &titulo)
	fmt.Print("Conteudo: ")
	fmt.Scanf("%s", &conteudo)
	art.Title = titulo
	art.Content = conteudo

	models.FazerArtigo(art)
	router := gin.Default()
	router.GET("/artigos", getArtigos)
	router.Run(":8080")
}
