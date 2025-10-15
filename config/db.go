package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "sqlserver://server=localhost;database=shop.go;user id=sa;password=My@Password123;encrypt=disable"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
