package main

import (
	"poker/api"
	"poker/api/kafka"
	"poker/api/oauth"
)

func main() {
	//grpc.RunGrpcSetver()
	go kafka.KafkaRead()
	oauth.OAuthChan = make(chan string, 1)
	api.RunRestServer()
}
