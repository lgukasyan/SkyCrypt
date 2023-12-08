package config

import (
	"github.com/joho/godotenv"
)

// Load environmental variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}
