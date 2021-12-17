/**
The bind route gets the users name
and binds it with "Hello {name}" message.

*/

package bind

import (
	"io"
	"net/http"
)

// Bind handles the "/bind" route
func Bind(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_, _ = io.WriteString(w, "Hello "+name)
}
