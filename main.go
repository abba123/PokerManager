package main

import (
	api "poker/apis"
)

func main() {

	//api.RunGrpcSetver()
	api.Tokens = map[string]bool{}
	api.RunRestServer()
}
