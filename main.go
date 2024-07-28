package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	internalHTTP "github.com/amirhnajafiz/letsgo/internal/http"
	"github.com/amirhnajafiz/letsgo/pkg/converter"
)

// system's variables
var (
	httpPort     int
	readTimeout  int
	writeTimeout int
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

func main() {
	// load env variables
	httpPort = converter.EnvToInt("GOHTTP_PORT", 8080)
	readTimeout = converter.EnvToInt("GOHTTP_READ_TO", 5)
	writeTimeout = converter.EnvToInt(("GOHTTP_WRITE_TO"), 10)

	// create a new net/http server instance
	app := setupServer()
	if app == nil {
		panic("failed to create a new application")
	}

	fmt.Printf("http server bounded to port  %d on this machine \n\tand its ready to handle input requests ...\n", httpPort)

	// start the http server
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
