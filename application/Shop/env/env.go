package env

import (
	"fmt"
	"os"
)

// Get returns an argument from environment or a fallback value if key not found
func Get(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(fmt.Sprintf("Could not find env for key %s. Returning default value %s", key, fallback))
		return fallback
	}
	return value
}
