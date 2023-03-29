package server

import (
	"doc-center-api/domain/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigRotas(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1/")
	{
		usuarios := main.Group("usuarios")
		{
			usuarios.GET("/", handlers.GetAllUsers)
			usuarios.GET("/:id", handlers.GetUserById)
			usuarios.POST("/", handlers.CreateUser)
			usuarios.PUT("/", handlers.UpdateUser)
			usuarios.DELETE("/:id", handlers.DeleteUser)
		}

		auth := main.Group("verify", handlers.)

		arquivos := main.Group("arquivos")
		{
			arquivos.GET("/", controllers.Listar)
			arquivos.GET("/:id", controllers.ObterPorId)
			arquivos.POST("/", controllers.Criar)
			arquivos.PUT("/", controllers.Atualizar)
			arquivos.DELETE("/:id", controllers.Apagar)
		}
	}
	return router
}
