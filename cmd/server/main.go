/**
This is the starter of the project.
It sets the server up with its configurations.
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"cmd/routes/home"
)

func main() {
	// Configuration of routes
	http.HandleFunc("/", home.Home)

	// Starting the server
	fmt.Println("Server is listening on 127.0.0.1:8080 ...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
