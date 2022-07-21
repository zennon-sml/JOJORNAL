package server

import(
    "github.com/gin-gonic/gin"
    "log"
    "github.com/zennon-sml/JOJORNAL/routes"
)

type Server struct{
    Porta string
    Servidor *gin.Engine
}

func NovoServidor() Server {
    return Server{
        Porta: "8080",
        Servidor: gin.Default(),
    }
}

func (s *Server) Rodar(){
   router := routes.ConfigurarRotas(s.Servidor)

   log.Print("Servidor Rodando Na Porta: ", s.Porta)
   log.Fatal(router.Run(":" + s.Porta))
}
