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
func GetPostgres() (*gorm.DB, error) {
	if PostgresDB != nil {
		return PostgresDB, nil
	}

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, databases)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		first_fiber.Host,
		first_fiber.User,
		first_fiber.Password,
		first_fiber.Db,
		first_fiber.Port,
	)

	cnf := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	var err error
	PostgresDB, err = gorm.Open(postgres.New(cnf), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := PostgresDB.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute)

	log.Info("POSTGRES LOADED")
	return PostgresDB, nil
}
