package main

import(
    //"github.com/zennon-sml/JOJORNAL/controller"
    "github.com/zennon-sml/JOJORNAL/server"
    "github.com/zennon-sml/JOJORNAL/database"
)

func main() {
    database.FazerCon()
    servidor := server.NovoServidor()
    servidor.Rodar()
}
