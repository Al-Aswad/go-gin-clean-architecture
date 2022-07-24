package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" autoincrement:"true" index:"true"`
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null;omitempty" json:"password"`
	Email     string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (e *User) TableName() string {
	return "users"
}
