package responses

// ObjectResponse is a struct used for returning user
// get request as a JSON object.
type ObjectResponse struct {
	Key    string `json:"key"`
	Values string `json:"value"`
}
