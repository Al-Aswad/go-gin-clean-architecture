package repositories

import (
	"gin-note-app/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	VerifyUser(email string, password string) interface{}
}
