package main

import (
	//"fmt"

	"fmt"
	"poker/poker"
)

func main() {
	//api.RunServer()
	//api.RunGrpcSetvet()
	tables := poker.Parse()
	for _, t := range tables {
		fmt.Println(t)
	}
}
