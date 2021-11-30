package main

import (
	//"fmt"

	api "poker/apis"
	"poker/poker"
)

func main() {
	//api.RunServer()
	api.RunGrpcSetvet()
	poker.Parse()
}
