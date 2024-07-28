package converter

import (
	"os"
	"strconv"
)

// EnvToInt function gets a name of an env variable and a defaultValue.
// Then, it will try to read the env variable, if an error occures, it returns the defaultValue.
func EnvToInt(name string, defaultValue int) int {
	if value, err := strconv.Atoi(os.Getenv(name)); err != nil {
		return defaultValue
	} else {
		return value
	}
}
