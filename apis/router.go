package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"poker/poker"
)

func RunRestServer() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/getwinrate/", getWinRate)
	router.GET("/", getHand)
	router.PUT("/", putHand)
	router.Run(":80")
}
