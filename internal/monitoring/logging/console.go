package logging

import "log"

// consoleLogr supports Logr interface. It displays the output to the process
// stdout (console).
type consoleLogr struct{}

func (c consoleLogr) Info(messages ...string) {
	format := "[INFO] %s"

	for _, msg := range messages {
		log.Printf(format, msg)
	}
}

func (c consoleLogr) Warning(messages ...string) {
	format := "[WARN] %s"

	for _, msg := range messages {
		log.Printf(format, msg)
	}
}

func (c consoleLogr) Error(err error, messages ...string) {
	format := "[ERROR] %s"

	log.Println(err)

	for _, msg := range messages {
		log.Printf(format, msg)
	}
}
