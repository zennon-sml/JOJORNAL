package controllers

import(
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/zennon-sml/JOJORNAL/database"
	"github.com/zennon-sml/JOJORNAL/models"
)

func NovoUsuarioForm(c *gin.Context){
	a := models.DadosPagina{Titulo: "Cadastrar Usuario"}
	c.HTML(200, "formUsuario.html", a) 
}

func NovoUsuario(c *gin.Context){
	var usuario models.Usuario
	if err := c.ShouldBind(&usuario); err != nil{
		fmt.Println("erro bindando: ", err)
	}
	bd := database.PegarBD()
	bd.Debug().Create(&usuario)
	dados := models.DadosPagina2{Titulo: "Usuario Criado",Usuario: usuario,}
	c.HTML(200, "jojornal.html", dados)
}

func EntrarForm(c *gin.Context){
	dados := models.DadosPagina2{Titulo: "Fazer Login"}
	c.HTML(200, "usuarioEntrarForm.html", dados)
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