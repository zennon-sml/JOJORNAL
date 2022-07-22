package routes

import(
    "github.com/gin-gonic/gin"
    "github.com/zennon-sml/JOJORNAL/controllers"
)

func ConfigurarRotas(router *gin.Engine) *gin.Engine{
    main := router.Group("jojornal/v1")
    {
        artigos := main.Group("artigos")
        {
            artigos.GET("/test", controllers.Test)
            artigos.Static("/static", "./static/")   //load static folde
            artigos.LoadHTMLGlob("templates/*.html") //load templates
            //routes
            artigos.GET("/artigos", models.PegarTodosArtigos)
            artigos.GET("/artigos/:id", models.PegarArtigo)
            artigos.GET("/artigos/novo", models.FormArtigo)
            artigos.POST("/artigos/novo", models.FazerArtigo)
            artigos.POST("/artigos/:id", models.ApagarArtigo)
            artigos.POST("/artigos/atualizar/form/:id", models.FormAtualizarArtigo)
            artigos.POST("/artigos/atualizar/:id", models.AtualizarArtigo)

        }
    }

    return router
}
