package logging

// Logr interface shows the system logging methods.
type Logr interface {
	Info(messages ...string)
	Warning(messages ...string)
	Error(err error, messages ...string)
}

// NewLogr returns a struct that supports Logr interface. In our case
// it returns a consoleLogr instance.
func NewLogr() Logr {
	return &consoleLogr{}
}
