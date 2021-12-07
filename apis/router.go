package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"poker/poker"
)

func RunRestServer() {
	router := gin.Default()

	router.Use(cors.New(CorsConfig()))
	router.POST("/", login)
	router.PUT("/", register)
	router.DELETE("/", middlewaree, logout)

	router.GET("/getwinrate/", middlewaree, getWinRate)

	router.GET("/hand/", middlewaree, getHand)
	router.PUT("/hand/", middlewaree, putHand)

	router.GET("/oauth/access/", googleOAuthAccess)
	router.GET("/oauth/login/", googleOAuthLogin)

	router.Run(":80")
}

// 定義 cors-config
func CorsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	return corsConf
}
