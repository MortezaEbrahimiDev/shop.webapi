package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Barcode     string  `gorm:"uniqueIndex"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	Description string
}
