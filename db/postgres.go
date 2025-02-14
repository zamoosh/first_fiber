package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
	user       = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASS")
	host       = os.Getenv("DB_HOST")
	db         = os.Getenv("DB_NAME")
	port       = os.Getenv("DB_PORT")
)

// InitMysql creates a connection to database
func initPostgres() (err error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		host,
		user,
		password,
		db,
		port,
	)

	cnf := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	PostgresDB, err = gorm.Open(postgres.New(cnf), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := PostgresDB.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
