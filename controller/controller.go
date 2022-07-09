package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/models"
)

var artigos []models.Artigo

func pegarArtigos(c *gin.Context) {
	c.JSON(http.StatusOK, artigos)
}

func main() {
	models.PegarTodosArtigos()
	// models.FazerArtigo(fazerBixo())
	router := gin.Default()
	router.GET("/artigos", mostrarArtigos)
	router.Run(":8080")
}

func fazerBixo() models.Artigo {
	reader := bufio.NewReader(os.Stdin)
	var art models.Artigo
	var lines []string
	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Print("Titulo: ")
		} else {
			fmt.Print("Conteudo: ")
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}
	art.Titulo = lines[0]
	art.Conteudo = lines[1]
	art.Criado = time.Now()
	return art
}

func criarArtigo(c *gin.Context) {
	var artigo models.Artigo
	c.BindJSON(&artigo)
}

func mostrarArtigos(c *gin.Context) {
	artigos := models.PegarTodosArtigos()
	c.BindJSON(&artigos)
}
