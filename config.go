package first_fiber

import (
	"os"

	"first_fiber/library/custom_log"
	"github.com/joho/godotenv"
)

var (
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresName     string
	PostgresPort     string
	MongoUser        string
	MongoPassword    string
	MongoHost        string
	MongoName        string
	MongoPort        string
)

func loadPostgres() {
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASS")
	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresName = os.Getenv("POSTGRES_NAME")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	custom_log.L.Info("PostgresDB CONFIGS LOADED")
}

func loadMongo() {
	PostgresUser = os.Getenv("MONGO_USER")
	PostgresPassword = os.Getenv("MONGO_PASS")
	PostgresHost = os.Getenv("MONGO_HOST")
	PostgresName = os.Getenv("MONGO_NAME")
	PostgresPort = os.Getenv("MONGO_PORT")
	custom_log.L.Info("MongoDB CONFIGS LOADED")
}

func loadLogger() {
	custom_log.L.Success("LOGGER LOADED")
	custom_log.L.Info("LOGGER LOADED")
}

func LoadConf() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	loadLogger()

	custom_log.L.Info("ENV LOADED")

	loadPostgres()
	loadMongo()

	return nil
}
