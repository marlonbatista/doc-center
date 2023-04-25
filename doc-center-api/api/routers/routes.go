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
		usuarios.GET("/:id", controllers.GetUserById)
		// usuarios.POST("/", controllers.CreateUser)
		// usuarios.PUT("/", controllers.UpdateUser)
		usuarios.DELETE("/:id", controllers.DeleteUser)
		usuarios.GET("/:id/permissoes", controllers.GetUserPermission)
	}

	login := routers.Group("login")
	{
		login.POST("/", controllers.Login)
	}

	arquivos := routers.Group("arquivos")
	{
		arquivos.GET("/", controllers.GetAllFiles)
		arquivos.GET("/:id", controllers.GetFileById)
		// arquivos.POST("/", controllers.CreateFile)
		// arquivos.PUT("/", controllers.UpdateFile)
		// arquivos.DELETE("/:id", controllers.DeleteFile)
	}

	permissoes := routers.Group("permissoes")
	{
		permissoes.GET("/", controllers.GetAllPermissions)
		permissoes.GET("/:id", controllers.GetPermissionById)
		permissoes.POST("/", controllers.CreatePermission)
		// permissao.PUT("/", controllers.UpdatePermission)
		// permissao.DELETE("/:id", controllers.DeletePermission)
	}

	dependentes := routers.Group("dependentes")
	{
		dependentes.GET("/", controllers.GetAllDependents)
		dependentes.GET("/:id", controllers.GetDependentById)
		// dependentes.POST("/", controllers.CreateDependent)
		// dependentes.PUT("/", controllers.UpdateDependent)
		// dependentes.DELETE("/:id", controllers.DeleteDependent)
	}

	return routers
}
