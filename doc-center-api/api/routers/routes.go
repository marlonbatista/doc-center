package routers

import (
	"doc-center-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRotas() *gin.Engine {
	routers := gin.Default()
	usuario := routers.Group("usuario")
	{
		usuario.GET("/", controllers.GetAllUsers)
		usuario.GET("/:id", controllers.GetUserById)
		// usuarios.POST("/", controllers.CreateUser)
		// usuarios.PUT("/", controllers.UpdateUser)
		usuario.DELETE("/:id", controllers.DeleteUser)
		usuario.GET("/:id/permissoes", controllers.GetUserPermission)
	}

	login := routers.Group("login")
	{
		login.POST("", controllers.Login)
	}

	arquivos := routers.Group("arquivos")
	{
		arquivos.GET("/", controllers.GetAllFiles)
		arquivos.GET("/:id", controllers.GetFileById)
		// arquivos.POST("/", controllers.CreateFile)
		// arquivos.PUT("/", controllers.UpdateFile)
		// arquivos.DELETE("/:id", controllers.DeleteFile)
	}

	permissao := routers.Group("permissao")
	{
		permissao.GET("/", controllers.GetAllPermissions)
		permissao.GET("/:id", controllers.GetPermissionById)
		// permissao.POST("/", controllers.CreatePermission)
		// permissao.PUT("/", controllers.UpdatePermission)
		// permissao.DELETE("/:id", controllers.DeletePermission)
	}

	dependente := routers.Group("dependente")
	{
		dependente.GET("/", controllers.GetAllDependents)
		dependente.GET("/:id", controllers.GetDependentById)
		// dependentes.POST("/", controllers.CreateDependent)
		// dependentes.PUT("/", controllers.UpdateDependent)
		// dependentes.DELETE("/:id", controllers.DeleteDependent)
	}

	return routers
}
