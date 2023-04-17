package main

import (
	"fmt"
	"pi-coursework-server/api"
)

func main() {
	fmt.Println("Start server")
	// api.
	// r := router

	r := api.SetupRouter()

	r.Run()

	// rou

}
