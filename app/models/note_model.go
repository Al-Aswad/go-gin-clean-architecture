package models

import "time"

type Note struct {
	ID          uint   `gorm:"primary_key" autoincrement:"true" index:"true"`
	Title       string `gorm:"not null"`
	Body        string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func (e *Note) TableName() string {
	return "notes"
}
