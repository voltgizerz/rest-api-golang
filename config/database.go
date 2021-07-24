package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() (bool) {
	dsn := GetHost()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
