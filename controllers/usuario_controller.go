package controllers

import(
	"fmt"
    "github.com/gin-gonic/gin"
	// "github.com/zennon-sml/JOJORNAL/database"
	"github.com/zennon-sml/JOJORNAL/models"
)

func NovoUsuarioForm(c *gin.Context){
	a := models.DadosPagina{Titulo: "titulo foda"}
	c.HTML(200, "formUsuario.html", a) 
}

func NovoUsuario(c *gin.Context){
	var usuario models.Usuario
	if err := c.ShouldBind(&usuario); err != nil{
		fmt.Println("erro bindando: ", err)
	}
	c.HTML(200, "usuarioEntrar.html", usuario)
}

func Entrar(c *gin.Context){
	var usuario models.Usuario
	if err := c.ShouldBind(&usuario); err != nil{
		fmt.Println("erro bindando: ", err)
	}
	dados := models.DadosPagina2{
		Titulo: "Novo Usuario",
		Usuario: usuario,
	}
	c.HTML(200, "usuario.html", dados)
}