package infrastructure

import (
	"os"
	"strconv"
)

// GetenvBool from env variable.
func GetenvBool(key string) bool {
	b, _ := strconv.ParseBool(os.Getenv(key))
	return b
}
