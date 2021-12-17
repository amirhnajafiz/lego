package server

import (
	"net/http"
	"strconv"
	"time"

	"cmd/routes/bind"
	"cmd/routes/home"
)

func SetupServer(port int) *http.Server {
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
