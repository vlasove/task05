package helper

import "os"

// EnvLoader ...
func EnvLoader(name, defaultEnvString string) string {
	p := os.Getenv(name)
	if len(p) == 0 {
		return defaultEnvString
	}
	return p
}
