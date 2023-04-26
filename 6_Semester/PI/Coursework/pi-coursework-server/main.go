package main

import (
	"fmt"
	"pi-coursework-server/api"
)

func main() {
	r := api.SetupRouter()

	fmt.Println("Start server")

	r.Run()
}
