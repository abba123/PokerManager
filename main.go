package main

import (
	api "poker/apis"
	"poker/apis/oauth"
)

func main() {
	//api.RunGrpcSetver()
	oauth.OAuthChan = make(chan string, 1)
	api.RunRestServer()

}
