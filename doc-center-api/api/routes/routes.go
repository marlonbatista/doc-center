package routes

import (
	"doc-center-api/api/controllers"
	"doc-center-api/api/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRotas() *gin.Engine{

	routers := gin.Default()
	
	// Rotas livres
	routers.POST("/login", controllers.Login)
	routers.POST("/signup", controllers.Signup)

	routers.Use(middleware.Authenticate())
	users := routers.Group("users")
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/", controllers.UpdateUser)
	}
	return routers
}
