/**
This is the starter of the project.
It sets the server up with its configurations.
*/

package main

import (
	"cmd/config/server"
	"fmt"
	"log"
	"os"
)

func main() {
	// Starting the server
	app := server.SetupServer(8080)
	if app == nil {
		os.Exit(-1)
	}

	fmt.Println("Server is listening on 127.0.0.1:8080 ...")
	err := app.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
