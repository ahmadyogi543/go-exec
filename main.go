package main

import (
	"github.com/ahmadyogi543/go-exec/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	v1 := app.Group("/api/v1")

	routes.Main(v1)
	routes.Compiler(v1)

	app.Run()
}
