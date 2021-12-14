package main

import (
	"poker/api"
	"poker/api/grpc"
	"poker/api/oauth"
)

func main() {
	grpc.RunGrpcSetver()
	oauth.OAuthChan = make(chan string, 1)
	api.RunRestServer()

}
