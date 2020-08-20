package helper

import (
	"fmt"
	"os"
	"strconv"
)

func GetStrFromEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}

func GetBoolFromEnv(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	result, err := strconv.ParseBool(val)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse env input: %s", key))
	}
	return result
}

func GetIntFromEnv(key string, defaultVal int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse env input: %s", key))
	}
	return result
}