package logging

// consoleLogr supports Logr interface. It displays the output to the process
// stdout (console).
type consoleLogr struct{}

func (c consoleLogr) Info(messages ...string) {

}

func (c consoleLogr) Warning(messages ...string) {

}

func (c consoleLogr) Error(err error, messages ...string) {

}
