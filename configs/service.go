package configs

import (
	"log"
	"os"
	"strconv"
)

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	res, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid port number: %v", err)
		return defaultValue
	}
	return res
}

func getString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
