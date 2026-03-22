package main

import (
	"amarolio-auth/src/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Run()
}
