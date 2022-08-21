/**
This is the starter of the project.
It sets the server up with its configurations.
*/

package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amirhnajafiz/lets-go/cmd/routes/bind"
	"github.com/amirhnajafiz/lets-go/cmd/routes/home"
)

// SetupServer function will create a http server for us
func setupServer(port int) *http.Server {
	// Configuration of routes
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/bind", bind.Bind)

	// Setting up the server
	server := http.Server{
		Addr:        ":" + strconv.Itoa(port),
		ReadTimeout: time.Second * 5,
	}

	return &server
}

func Execute() {
	// Starting the server
	app := setupServer(8080)
	if app == nil {
		os.Exit(-1)
	}

	fmt.Println("Server is listening on 127.0.0.1:8080 ...")
	err := app.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
