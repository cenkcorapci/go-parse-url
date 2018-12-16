package utils

import "os"

func GetEnv(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
