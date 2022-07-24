package models

import "time"

type User struct {
	ID        uint64     `gorm:"primary_key" autoincrement:"true" index:"true" json:"id"`
	Username  string     `gorm:"not null;unique" json:"username"`
	Password  string     `gorm:"not null;omitempty" json:"-"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Token     string     `gorm:"not null" json:"token"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (e *User) TableName() string {
	return "users"
}
