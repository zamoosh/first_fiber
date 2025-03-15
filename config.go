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
	MongoUser = os.Getenv("MONGO_USER")
	MongoPassword = os.Getenv("MONGO_PASS")
	MongoHost = os.Getenv("MONGO_HOST")
	MongoName = os.Getenv("MONGO_NAME")
	MongoPort = os.Getenv("MONGO_PORT")
	custom_log.L.Info("MongoDB CONFIGS LOADED")
}

func loadLogger() {
	custom_log.Default()
	custom_log.L.Info("LOGGER LOADED")
}

func LoadConf() error {
	loadLogger()

	err := godotenv.Load()
	if err != nil {
		return err
	}
	custom_log.L.Info("ENV LOADED")

	loadPostgres()
	loadMongo()

	return nil
}
