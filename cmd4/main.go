package main

import (
	"os"

	"first_fiber/library/utils/auth"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token, err := auth.GenerateToken(10, auth.RefreshToken)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info(token)

	valid, err := auth.VerifyToken(token)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if valid {
		log.Info("token is valid")
	}
}
