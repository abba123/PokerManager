package main

import (
	api "poker/apis"
	oauth "poker/apis/OAuth"
)

func main() {	
	//api.RunGrpcSetver()
	oauth.OAuthChan = make(chan string,1)
	api.RunRestServer()
	//url := oauth.GenerateURL()
	//fmt.Println(url)
}
