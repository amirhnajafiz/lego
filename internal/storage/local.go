package storage

// localstorage is a struct that supports Storage interface. It uses Golang
// map object to store the system's data.
type localstorage struct {
	mapStore map[string]*interface{}
}

// Store an object with specific key.
func (l localstorage) Store(key string, object *interface{}) {
	l.mapStore[key] = object
}

// Receive an object by its key.
func (l localstorage) Receive(key string) *interface{} {
	if value, ok := l.mapStore[key]; ok {
		return value
	}

	return nil
}

// Delete an object by its key.
func (l localstorage) Delete(key string) {
	delete(l.mapStore, key)
}
