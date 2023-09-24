package util

import (
	"os"
)

func ReadEnvs(envKeys ...string) string {
	for _, envKey := range envKeys {
		value := os.Getenv(envKey)
		if value != "" {
			return value
		}
	}

	return ""
}
