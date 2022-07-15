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
	router.GET("/artigos/novo", models.FormArtigo)
	router.POST("/artigos/novo", models.FazerArtigo)
	router.POST("/artigos/:id", models.ApagarArtigo)
	// router.PUT("/artigos/:id", models.AtualizarArtigo)
	router.Run(":8080")
}
