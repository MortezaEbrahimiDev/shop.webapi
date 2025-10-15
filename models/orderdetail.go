package models

type OrderDetail struct {
	ID        uint    `gorm:"primary_key;AUTO_INCREMENT"`
	OrderID   uint    `gorm:"not null"`
	ProductId uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"not null"`
}
