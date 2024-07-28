package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	internalHTTP "github.com/amirhnajafiz/letsgo/internal/http"
)

// system's variables
var (
	httpPort     = 8080
	readTimeout  = 10
	writeTimeout = 5
)

// setupServer creates a net/http server and sets
// the internal router for user requests.
func setupServer() *http.Server {
	// create a new mux server
	muxServer := internalHTTP.Bootstrap()

	// setting up the http server
	server := http.Server{
		Addr:         ":" + strconv.Itoa(httpPort),
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
		Handler:      muxServer,
	}

	return &server
}

// loadEnvVariables read environmental variables values to update system's
// base variables.
func loadEnvVariables() {
	httpPort, _ = strconv.Atoi(os.Getenv("GOHTTP_PORT"))
	readTimeout, _ = strconv.Atoi(os.Getenv("GOHTTP_READ_TO"))
	writeTimeout, _ = strconv.Atoi(os.Getenv("GOHTTP_WRITE_TO"))
}

func main() {
	// load env variables
	loadEnvVariables()

	// create a new net/http server instance
	app := setupServer()
	if app == nil {
		panic("failed to create a new application")
	}

	fmt.Printf("http server bounded to port  %d on this machine \n\tand its ready to handle input requests ...", httpPort)

	// start the http server
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
