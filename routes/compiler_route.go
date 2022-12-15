package routes

import (
	"github.com/ahmadyogi543/go-exec/controllers"
	"github.com/gin-gonic/gin"
)

func Compiler(route *gin.RouterGroup) {
	root := route.Group("/compiler")
	{
		root.GET("/", controllers.MainCompiler)
		root.POST("/execute", controllers.ExecuteCompiler)
		root.GET("/list", controllers.ListCompiler)
		root.GET("/version", controllers.VersionCompiler)
	}
}
