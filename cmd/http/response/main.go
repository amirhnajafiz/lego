package response

import "time"

type BindResponse struct {
	Message string
	Time    time.Time
}

type HomeResponse struct {
	Status  string
	Message string
}
