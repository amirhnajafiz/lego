package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	internalHTTP "github.com/amirhnajafiz/letsgo/internal/http"
)

// setupServer creates a net/http server and sets
// the internal router for user requests.
func setupServer(port int) *http.Server {
	// create a new mux server
	muxServer := internalHTTP.Bootstrap()

	// setting up the http server
	server := http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		Handler:      muxServer,
	}

	return &server
}

func main() {
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
