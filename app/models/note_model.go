package models

import "time"

type Note struct {
	ID          uint       `gorm:"primary_key" autoincrement:"true" index:"true"`
	Title       string     `gorm:"not null" json:"title"`
	Body        string     `gorm:"not null" json:"body"`
	Description string     `gorm:"not null" json:"description"`
	UserID      uint64     `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

func (e *Note) TableName() string {
	return "notes"
}
