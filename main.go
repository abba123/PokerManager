package main

import (
	"poker/api"
	"poker/api/kafka"
	"poker/api/model"
	"poker/api/oauth"

	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()
	if viper.GetString("BACKEND") == ""{
		viper.Set("BACKEND", "127.0.0.1")
	}
	if viper.GetString("DATABASE") == "" {
		viper.Set("DATABASE", "127.0.0.1")
	}
	if viper.GetString("REDIS") == "" {
		viper.Set("REDIS", "127.0.0.1")
	}
	if viper.GetString("KAFKA") == "" {
		viper.Set("KAFKA", "127.0.0.1")
	}

	//grpc.RunGrpcSetver()
	go kafka.KafkaRead()
	oauth.OAuthChan = make(chan string, 1)
	model.InitDB("pokerdb")
	api.RunRestServer()

}
