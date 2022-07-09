package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

func main() {
	// models.FazerArtigo(fazerBixo())

	router := gin.Default()
	//routes
	router.GET("/artigos", models.PegarTodosArtigos)
	router.GET("/artigos/:id", models.PegarArtigo)
	router.DELETE("/artigos/:id", models.ApagarArtigo)
	router.PUT("/artigos", models.AtualizarArtigo)
	router.POST("/artigos", models.FazerArtigo)
	router.Run(":8080")
}
