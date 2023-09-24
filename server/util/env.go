package util

import (
	"fmt"
	"os"
)

func ReadEnvs(envKeys ...string) string {
	for _, envKey := range envKeys {
		value := os.Getenv(envKey)
		if value != "" {
			fmt.Print("ReadEnvs: " + envKey + " = " + value + "\n")
			return value
		}
	}

	return ""
}