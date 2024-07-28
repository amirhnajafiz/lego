package storage

// localStorage is a struct that supports Storage interface. It uses Golang
// map object to store the system's data.
type localStorage struct {
	mapStore map[string]*interface{}
}

// Store an object with specific key.
func (l *localStorage) Store(key string, object *interface{}) {
	l.mapStore[key] = object
}

// Receive an object by its key.
func (l *localStorage) Receive(key string) *interface{} {
	if value, ok := l.mapStore[key]; ok {
		return value
	}

	return nil
}

// Delete an object by its key.
func (l *localStorage) Delete(key string) {
	delete(l.mapStore, key)
}
