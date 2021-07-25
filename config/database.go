package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() (bool) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	dsn := GetHost()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
			log.Println(err)
	}
	DB = db
	return true
}

func GetHost() (string) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	host, _ := os.LookupEnv("DB_HOST_LOCAL")

	env, _ := os.LookupEnv("ENV")
	if env == "production" {
		host, _ = os.LookupEnv("DB_HOST_PRODUCTION")
	}
	return host
}
