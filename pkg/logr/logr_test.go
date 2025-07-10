package logr

import (
	"log"
	"testing"
)

// TestLogr tests the logger functionality.
// It creates a logger with debug level and logs messages at different levels.
func TestLogr(t *testing.T) {
	logger, err := New(Config{Level: "debug"})
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
