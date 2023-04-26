package main

import (
	"flag"
	"pi-coursework-client/controller"
)

func main() {
	var usernameFlag string
	var passwordFlag string
	flag.StringVar(&usernameFlag, "username", "PiDB", "user name")
	flag.StringVar(&passwordFlag, "password", "PiDB", "user's password")
	flag.Parse()

	// if err := controller.Auth(usernameFlag, passwordFlag); err != nil {
	// 	fmt.Println("Failed to auth")
	// 	return
	// }

	controller.EndlessInput()
}
