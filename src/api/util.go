package api

import (
	"os"
)

// GetEnv is utility function to get default environment variable if not defined
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
