package models

type Role struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"uniqueIndex;not null"`
	Users []User `gorm:"many2many:user_roles"`
}
