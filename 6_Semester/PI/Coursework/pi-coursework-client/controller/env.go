package controller

import (
	"os"
)

func GetEnvDefault(key string, defaultVal string) (value string) {
	value = defaultVal

	looked, ok := os.LookupEnv(key)
	if ok {
		value = looked
	}

	return
}
