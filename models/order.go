package models

import "time"

type Order struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT"`
	CustomerID uint   `gorm:"not null"`
	Customer   User   `gorm:"foreignkey:CustomerID"`
	Status     string `gorm:"not null"`
	CreateAt   time.Time
	Details    []OrderDetail
}
