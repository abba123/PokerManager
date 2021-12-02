package main

import (
	api "poker/apis"
)

func main() {

	//api.RunGrpcSetver()
	api.Token = map[string]bool{}
	api.RunRestServer()

}
