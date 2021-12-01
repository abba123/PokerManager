package main

import (
	api "poker/apis"
)

func main() {

	//api.RunGrpcSetver()
	api.Token = map[string]string{}
	api.RunRestServer()

}
