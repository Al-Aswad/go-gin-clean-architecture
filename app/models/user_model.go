package models

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key" autoincrement:"true" index:"true"`
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null;omitempty" json:"password"`
	Email     string `gorm:"not null;unique"`
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (e *User) TableName() string {
	return "users"
}
