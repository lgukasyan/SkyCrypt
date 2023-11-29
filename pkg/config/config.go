package config

import (
	"log"
	"github.com/joho/godotenv"
)

// Load environmental variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
		return
	}
}
