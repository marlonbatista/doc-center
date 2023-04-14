package routers

import (
	"doc-center-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRotas() *gin.Engine {
	routers := gin.Default()
	usuarios := routers.Group("usuarios")
	{
		usuarios.GET("/", controllers.GetAllUsers)
		// usuarios.GET("/:id", controllers.GetUserById)
		// usuarios.POST("/", controllers.CreateUser)
		// usuarios.PUT("/", controllers.UpdateUser)
		// usuarios.DELETE("/:id", controllers.DeleteUser)
	}

	login := routers.Group("login")
	{
		login.POST("/", controllers.Login)
	}

	return routers
}
