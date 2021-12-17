/**
The home route returns a json response
with status OK and message home.

*/

package home

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Home The home-response function handles the '/' route
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	var data = Response{}
	data.Status = strconv.Itoa(http.StatusAccepted)
	data.Message = "Welcome Home"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(data)
}
