package server

import(
    "github.com/gin-gonic/gin"
    "log"
    "github.com/zennon-sml/JOJORNAL/routes"
    "os"
)

type Server struct{
    Porta string
    Servidor *gin.Engine
}

func NovoServidor() Server {
    port := os.Getenv("PORT")
    return Server{
        Porta: port,
        Servidor: gin.Default(),
    }
}

func (s *Server) Rodar(){
   router := routes.ConfigurarRotas(s.Servidor)

   log.Print("Servidor Rodando Na Porta: ", s.Porta)
   log.Fatal(router.Run(":" + s.Porta))
}
