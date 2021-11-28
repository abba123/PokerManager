package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"poker/poker"
)

func RunServer() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", get)

	router.Run(":80")
}
