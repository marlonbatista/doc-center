package routes

import (
	"doc-center-api/api/controllers"
	"doc-center-api/api/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigRotas() *gin.Engine {

	routers := gin.Default()

	// Rotas livres
	routers.POST("/login", controllers.Login)
	routers.POST("/signup", controllers.Signup)

	routers.Use(middleware.Authenticate())
	users := routers.Group("users")
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.GET("/:id/permissions", controllers.GetUserPermission)
		users.GET("name/:name", controllers.GetUserByName)
		users.PUT("/", controllers.UpdateUser)
	}
	files := routers.Group("files")
	{
		files.GET("/", controllers.GetAllFiles)
		files.GET("/:id", controllers.GetFileById)
		files.POST("/", controllers.CreateFile)
		files.PUT("/", controllers.UpdateFile)
		files.DELETE("/:id", controllers.DeleteFile)
	}
	permissions := routers.Group("permissions")
	{
		permissions.GET("/", controllers.GetAllPermissions)
		permissions.GET("owner/:id", controllers.GetOwnerPermission)
		permissions.GET("guest/:id", controllers.GetGuestPermission)
		permissions.POST("/", controllers.CreatePermission)
		permissions.PUT("/", controllers.UpdatePermission)
		permissions.DELETE("/:id", controllers.DeletePermission)
	}
	dependents := routers.Group("dependents")
	{
		dependents.GET("/", controllers.GetAllDependents)
		dependents.GET("/:id", controllers.GetDependentById)
		// dependents.POST("/", controllers.CreateDependent)
		// dependents.PUT("/", controllers.UpdateDependent)
		// dependents.DELETE("/:id", controllers.DeleteDependent)
	}
	return routers
}
