package services

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/models"
)

type UserService interface {
	Create(user dto.RegisterUserDto) (models.User, error)
	// Create(user dto.RegisterUserDto) (dto.RegisterUserDto, error)
	// Create(person *model.Person) (*model.Person, error)
}
