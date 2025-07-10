package uuid

import (
	"github.com/google/uuid"
)

// New generates a new UUID.
func New() uuid.UUID {
	return uuid.New()
}

// NewString generates a new UUID as a string.
func NewString() string {
	return uuid.NewString()
}
