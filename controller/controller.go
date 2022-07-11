package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

func main() {
	// models.FazerArtigo(fazerBixo())

	router := gin.Default()
	router.Static("/static", "./static/")
	router.LoadHTMLGlob("templates/*.html") //templates
	//routes
	router.GET("/artigos", models.PegarTodosArtigos)
	router.GET("/artigos/:id", models.PegarArtigo)
	router.DELETE("/artigos/:id", models.ApagarArtigo)
	router.PUT("/artigos", models.AtualizarArtigo)
	router.POST("/artigos", models.FazerArtigo)
	router.GET("/", models.HTMLExemplo)
	router.Run(":8080")
}
