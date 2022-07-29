package repositories

import (
	"go-gin-note-app/app/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	VerifyUser(email string, password string) interface{}
}
