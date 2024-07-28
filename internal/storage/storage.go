package storage

// Storage interface shows the storage methods that are used
// to perform memory operations.
type Storage interface {
	Store(key string, object *interface{})
	Receive(key string) *interface{}
	Delete(key string)
}

// NewStorage returns a struct which supports Storage interface, in our case
// it returns a localstorage instance.
func NewStorage() Storage {
	return &localstorage{
		mapStore: make(map[string]*interface{}),
	}
}
