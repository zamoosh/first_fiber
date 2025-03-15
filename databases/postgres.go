package databases

import (
	"fmt"
	"time"

	"first_fiber"

	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
)

// GetPostgres creates a connection to Postgres Database.
func GetPostgres() *gorm.DB {
	if PostgresDB != nil {
		return PostgresDB
	}

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, databases)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		first_fiber.PostgresHost,
		first_fiber.PostgresUser,
		first_fiber.PostgresPassword,
		first_fiber.PostgresName,
		first_fiber.PostgresPort,
	)

	cnf := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	var err error
	PostgresDB, err = gorm.Open(postgres.New(cnf))
	if err != nil {
		log.Fatalf("Could not open connection. %s", err)
	}

	sqlDB, err := PostgresDB.DB()

	if err != nil {
		log.Fatalf("Could not get sqlDB. %s", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute)

	log.Info("POSTGRES LOADED")
	return PostgresDB
}
