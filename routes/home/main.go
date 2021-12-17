/**
The home route returns a json response
with status OK and message home.

*/

package home

import (
	"io"
	"net/http"
)

// Home The home function handles the '/' route
func Home(w http.ResponseWriter, r *http.Request) {
	message := "Server response"
	if r.Method != "GET" {
		message = "Bad request"
	}
	_, _ = io.WriteString(w, message)
}
