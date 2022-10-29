package config

import (
	"log"
	"os"
	"time"

	"rest-api-golang/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func InitDB() Database {
	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
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
		logger.Log.Error(err)
	}

	return Database{DB: db}
}

func GetHost() string {
	if err := godotenv.Load(); err != nil {
		logger.Log.Error("No .env file found")
	}

	host, _ := os.LookupEnv("DB_HOST_LOCAL")

	env, _ := os.LookupEnv("ENV")
	if env == "production" {
		host, _ = os.LookupEnv("DB_HOST_PRODUCTION")
	}
	return host
}
