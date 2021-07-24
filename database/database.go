package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func Database() (*gorm.DB) {
	dsn := "root:@tcp(localhost)/uncle_loe?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
			panic(err)
	}
	log.Println("DB Connected")
	return db
}