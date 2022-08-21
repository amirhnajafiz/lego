package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amirhnajafiz/lets-go/cmd/http/response"
)

// Bind handles the "/bind" route
func Bind(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	data := response.BindResponse{}
	data.Message = "Hello " + name
	data.Time = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(data)
}

// Home The home-response function handles the '/' route
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	var data = response.HomeResponse{}
	data.Status = strconv.Itoa(http.StatusAccepted)
	data.Message = "Welcome Home"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(data)
}
