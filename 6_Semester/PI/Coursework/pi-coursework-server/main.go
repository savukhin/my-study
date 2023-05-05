package main

import (
	"fmt"
	"pi-coursework-server/api"
	mainexecutor "pi-coursework-server/main_executor"
	"pi-coursework-server/table"
)

func main() {
	fmt.Println("TABLE PATH IS ", table.TABLES_PATH)
	err := mainexecutor.LoadInitialState()
	if err != nil {
		panic(err)
	}

	r := api.SetupRouter()

	fmt.Println("Start server")

	r.Run()
}
