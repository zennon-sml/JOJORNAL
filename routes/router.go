package routes

import(
    "github.com/gin-gonic/gin"
    "github.com/zennon-sml/JOJORNAL/controllers"
)

func ConfigurarRotas(router *gin.Engine) *gin.Engine{
    router.Static("/static", "./static/")
    router.LoadHTMLGlob("templates/*.html")
    main := router.Group("jojornal/v1")
    {
        artigos := main.Group("artigos")
        {
            artigos.GET("/test", controllers.Test)
            artigos.GET("/", controllers.PegarTodosArtigos)
            artigos.GET("/:id", controllers.PegarArtigo)
            artigos.POST("/:id", controllers.ApagarArtigo)
            artigos.GET("/novo", controllers.FormArtigo)
            artigos.POST("/novo", controllers.FazerArtigo)
            artigos.POST("/atualizar/form/:id", controllers.FormAtualizarArtigo)
            artigos.POST("/atualizar/:id", controllers.AtualizarArtigo)
        }
    }

    return router
}
