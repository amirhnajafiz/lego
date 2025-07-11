package configs

import (
	"os"
	"testing"
)

// constant content
const (
	content = `a: "test"
b: 42
c: true`
)

// exampleConfig is a struct for testing purposes.
type exampleConfig struct {
	A string `koanf:"a"`
	B int    `koanf:"b"`
	C bool   `koanf:"c"`
	D string `koanf:"d"`
}

// createTestFile creates a test file with the given content.
func createTestFile(filePath, content string) error {
	// create a new file with the given content
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// deleteTestFile deletes the test file if it exists.
func deleteTestFile(filePath string) error {
	// delete the test file if it exists
	if _, err := os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	}
	return nil
}

// TestConfigs tests the configuration loading functionality.
// It creates a temporary YAML file, loads it using the New function,
// and checks if the values are correctly parsed into the exampleConfig struct.
func TestConfigs(t *testing.T) {
	filePath := "test_config.yaml"
	defer deleteTestFile(filePath)

	// create a new yaml file with content
	err := createTestFile(filePath, content)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// set the environment variable for testing
	os.Setenv("MBE_d", "env value")

	// create an instance of exampleConfig
	instance := &exampleConfig{}

	// load the configuration
	cfg, err := New(filePath, "MBE_", instance)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// assert the values
	if cfg.(*exampleConfig).A != "test" {
		t.Errorf("expected A to be 'test', got '%s'", cfg.(*exampleConfig).A)
	}
	if cfg.(*exampleConfig).B != 42 {
		t.Errorf("expected B to be 42, got %d", cfg.(*exampleConfig).B)
	}
	if cfg.(*exampleConfig).C != true {
		t.Errorf("expected C to be true, got %v", cfg.(*exampleConfig).C)
	}
	if cfg.(*exampleConfig).D != "env value" {
		t.Errorf("expected D to be 'env value', got '%s'", cfg.(*exampleConfig).D)
	}
}
