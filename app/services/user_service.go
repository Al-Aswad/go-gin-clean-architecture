package services

import (
	"go-gin-note-app/app/dto"
	"go-gin-note-app/app/models"
)

type UserService interface {
	Create(user dto.RegisterUserDto) (models.User, error)
	// Create(user dto.RegisterUserDto) (dto.RegisterUserDto, error)
	// Create(person *model.Person) (*model.Person, error)
}
