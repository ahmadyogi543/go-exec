package controllers

import "github.com/gin-gonic/gin"

func MainMain(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to go-exec REST API",
	})
}

func PingMain(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
