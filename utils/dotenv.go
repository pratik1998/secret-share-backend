package util

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	env := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}

func GodotEnvInt(key string) int {
	env := make(chan int, 1)

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(".env")
		value, _ := strconv.Atoi(os.Getenv(key))
		env <- value
	} else {
		value, _ := strconv.Atoi(os.Getenv(key))
		env <- value
	}

	return <-env
}
