package api

import (
	"log"
	"os"
	"strconv"
)

// GetEnv is utility function to get default environment variable if not defined
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetIntEnv is utility function to get default int environment variable if not defined
func GetIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		ret, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		return ret
	}
	return fallback
}
