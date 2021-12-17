/**
This is the starter of the project.
It sets the server up with its configurations.
*/

package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Configuration of routes
	http.HandleFunc("/", handler)

	// Starting the server
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Server is listening on 127.0.0.1:8080 ...")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Server response"
	if r.Method != "get" {
		message = "Bad request"
	}
	_, _ = io.WriteString(w, message)
}