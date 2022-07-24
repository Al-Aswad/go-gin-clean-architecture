package repositories

import (
	"go-gin-clean-architecture/app/models"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
}
