package env

import (
	"fmt"
	"os"
	"strconv"
)

// String defines a string environment variable with specified name, and default value.
// The return value is a string variable that stores the value of the environment variable.
func String(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok && value != "" {
		return value
	}
	return defaultValue
}

// Bool defines a bool environment variable with specified name, and default value.
// The return value is a bool variable that stores the value of the environment variable.
func Bool(name string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(name); ok {
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			panic(fmt.Sprintf("%s value should be true or false, got %s", name, value))
		}
		return boolVal
	}

	return defaultValue
}
