package converter

import "encoding/json"

// StructToJSON gets a Golang interface and returns an array of bytes
// which can be converted to a JSON object (as string).
func StructToJSON(obj interface{}) ([]byte, error) {
	if arrayOfBytes, err := json.Marshal(obj); err != nil {
		return nil, err
	} else {
		return arrayOfBytes, nil
	}
}
