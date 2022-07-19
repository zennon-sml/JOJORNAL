package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

func Main() {
	// models.FazerArtigo(fazerBixo())

	router := gin.Default()
	router.Static("/static", "./static/")   //load static folder
	router.LoadHTMLGlob("templates/*.html") //load templates
	//routes
	router.GET("/artigos", models.PegarTodosArtigos)
	router.GET("/artigos/:id", models.PegarArtigo)
	router.GET("/artigos/novo", models.FormArtigo)
	router.POST("/artigos/novo", models.FazerArtigo)
	router.POST("/artigos/:id", models.ApagarArtigo)
	router.POST("/artigos/atualizar/form/:id", models.FormAtualizarArtigo)
	router.POST("/artigos/atualizar/:id", models.AtualizarArtigo)

	router.Run(":8080")
}
