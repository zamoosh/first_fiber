package main

import (
	"os"

	"first_fiber/library/utils"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token, err := utils.GenerateToken(10, utils.RefreshToken)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info(token)

	valid, err := utils.VerifyToken(token)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if valid {
		log.Info("token is valid")
	}
}
