package models

type User struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}
