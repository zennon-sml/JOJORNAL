package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

var artigos []models.Artigo

func getArtigos(c *gin.Context) {
	c.JSON(http.StatusOK, artigos)
}

func main() {
	var artigo1 models.Artigo
	artigo1.ID = 1
	artigo1.Title = "JOJO"
	artigo1.Content = "DIIIOOOOOOO"

	a := models.Artigo{ID: 2, Title: "One Piece", Content: "skypie"}
	b := models.Artigo{ID: 3, Title: "NARUTO", Content: "BRIDGE"}
	artigos = append(artigos, artigo1, a, b)
	router := gin.Default()
	router.GET("/artigos", getArtigos)
	router.Run(":8080")
}
