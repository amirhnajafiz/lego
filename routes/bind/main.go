/**
The bind route gets the users name
and binds it with "Hello {name}" message.

*/

package bind

import (
	"encoding/json"
	"net/http"
	"time"
)

// Bind handles the "/bind" route
func Bind(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	data := Response{}
	data.Message = "Hello " + name
	data.Time = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(data)
}
