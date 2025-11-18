package util

import (
	"os"
)

// Gets int env variable, if not set, returns defaultVal.
func GetOptionalStringVariable(name string, defaultVal string) string {
	valueStr := os.Getenv(name)
	if valueStr != "" {
		return valueStr
	}

	return defaultVal
}
