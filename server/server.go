package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/routes"
)

type Server struct {
	Porta    string
	Servidor *gin.Engine
}

func NovoServidor() Server {
	//pegar porta do ambiente pra fazer deploy heroku
	port := os.Getenv("PORT")
	return Server{
		Porta:    port,
		Servidor: gin.Default(),
	}
}

func (s *Server) Rodar() {
	router := routes.ConfigurarRotas(s.Servidor)

	log.Print("Servidor Rodando Na Porta: ", s.Porta)
	log.Fatal(router.Run(":" + s.Porta))
}
