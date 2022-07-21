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
        }
    }

    return router
}
