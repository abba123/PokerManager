package api

import (
	"github.com/gin-gonic/gin"
	//"poker/poker"
)

func RunServer() {
	router := gin.Default()
	router.GET("/", get)

	router.Run(":80")
}