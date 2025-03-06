package first_fiber

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var (
	User       string
	Password   string
	Host       string
	Db         string
	Port       string
)

func LoadConf() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	log.Info("ENV LOADED")

	User = os.Getenv("POSTGRES_USER")
	Password = os.Getenv("POSTGRES_PASS")
	Host = os.Getenv("POSTGRES_HOST")
	Db = os.Getenv("POSTGRES_NAME")
	Port = os.Getenv("POSTGRES_PORT")
	log.Info("DB CONFIGS LOADED")

	return nil
}
