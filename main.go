package main


import(
	"poker/apis"
)

func main() {
	api.RunRestServer()
	//api.RunGrpcSetvet()
	/*
		tables := poker.Parse()
		for _, t := range tables {
			fmt.Println(t)
		}
	*/
}
