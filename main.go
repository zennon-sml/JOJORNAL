package main

import(
    //"github.com/zennon-sml/JOJORNAL/controller"
    "github.com/zennon-sml/JOJORNAL/server"
)

func main() {
//	controller.Main()
    servidor := server.NovoServidor()
    servidor.Rodar()
}
