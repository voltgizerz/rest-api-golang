package config

import (
	"log"
	"os"
	"time"

	"rest-api-golang/logger"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Database - .
type Database struct {
	DB *gorm.DB
}

// InitDB - .
func InitDB() Database {
	log := logger.Log.WithFields(logrus.Fields{
		"LstdFlags": log.LstdFlags,
	})

	newLogger := gormLogger.New(
		log, // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)

	dsn := GetHost()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Log.Fatal(err)
	}

	return Database{
		DB: db,
	}
}

func GetHost() string {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("No .env file found")
	}

	host, _ := os.LookupEnv("DB_HOST_LOCAL")

	env, _ := os.LookupEnv("ENV")
	if env == "production" {
		host, _ = os.LookupEnv("DB_HOST_PRODUCTION")
	}
	return host
}
