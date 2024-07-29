package v1

// objectResponse is a struct used for returning user
// get request as a JSON object.
type objectResponse struct {
	Key    string `json:"key"`
	Values string `json:"value"`
}
