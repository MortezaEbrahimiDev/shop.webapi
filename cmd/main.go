package cmd

import (
	"shop.go/config"
	"shop.go/models"
)

func main() {
	config.ConnectDb()

	config.DB.AutoMigrate(&models.User{},
		&models.Role{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Product{})
}
