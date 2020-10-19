package utils

import (
	"log"
	"os"
	"strconv"
)

// MustGet will return env or panic
func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Panicln("ENV missing, key: " + v)
	}
	return v
}

// MustGetBool will return env as bool or panic
func MustGetBool(key string) bool {
	v := MustGet(key)
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicln("ENV err: [" + v + "\n" + err.Error())
	}

	return b
}
