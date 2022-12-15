package controllers

import (
	"net/http"
	"time"

	"github.com/ahmadyogi543/go-exec/constants"
	"github.com/ahmadyogi543/go-exec/types"
	"github.com/ahmadyogi543/go-exec/utils"
	"github.com/gin-gonic/gin"
)

func MainCompiler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome, you can use /compiler/execute endpoint to execute your code",
	})
}

func ExecuteCompiler(ctx *gin.Context) {
	var mainErr error
	var codeRequest types.CodeRequest
	ctx.BindJSON(&codeRequest)

	extension := constants.COMPILERS[codeRequest.Language].Extension
	executable := constants.COMPILERS[codeRequest.Language].Executable

	codePath, err := utils.CreateCode(codeRequest.Code, extension)
	if err != nil {
		mainErr = err
	}

	codeOutput, _ := utils.ExecuteCode(codePath, codeRequest.Input, executable)
	// if err != nil {
	// 	mainErr = err
	// }

	_ = utils.RemoveFile(codePath)
	// if err != nil {
	// 	mainErr = err
	// }

	if mainErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "internal server error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"output":    codeOutput,
			"language":  codeRequest.Language,
			"timestamp": time.Now(),
		})
	}
}

func ListCompiler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, constants.COMPILERS)
}

func VersionCompiler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": "8.1.13",
	})
}
