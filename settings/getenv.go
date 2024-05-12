package settings

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		port = "3001"
	}
	return port
}
