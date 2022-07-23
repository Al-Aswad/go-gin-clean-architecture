package services

import "go-gin-clean-architecture/app/models"

type UserService interface {
	Create(user *models.User) (*models.User, error)
	// Create(person *model.Person) (*model.Person, error)
}
