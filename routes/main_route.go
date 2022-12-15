package routes

import (
	"github.com/ahmadyogi543/go-exec/controllers"
	"github.com/gin-gonic/gin"
)

func Main(route *gin.RouterGroup) {
	main := route.Group("/")
	{
		main.GET("/", controllers.MainMain)

		main.GET("/ping", controllers.PingMain)
	}
}
