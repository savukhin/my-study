package main

import (
	"flag"
	"pi-coursework-client/controller"
)

/*

BEGIN;
    CREATE TABLE musicians(name, surname, band);
    INSERT INTO musicians(surname, band, name) values (Shinoda, LinkinPark, Mike);
    INSERT INTO musicians(name, surname, band) values (Chester, Bennington, LinkinPark);
    INSERT INTO musicians(name, surname, band) values (Maybe, Baby, MaybeBaby);
COMMIT;
SELECT * FROM musicians;

*/

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
