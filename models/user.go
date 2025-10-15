package models

type User struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username     string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Roles        []Role `gorm:"many2many:user_roles"`
	Orders       []Order
}
