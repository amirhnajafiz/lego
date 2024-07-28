package measure

import (
	"unsafe"
)

// GetObjectSize returns the actual size of an interface in float format.
func GetObjectSize(obj interface{}) float64 {
	return float64(unsafe.Sizeof(obj))
}
