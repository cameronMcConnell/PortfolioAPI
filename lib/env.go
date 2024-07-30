package lib

import "os"

func ReadEnv(envVar string) string {
	return os.Getenv(envVar)
}