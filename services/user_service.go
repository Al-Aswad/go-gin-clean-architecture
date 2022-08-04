package services

import (
	"gin-note-app/dto"
	"gin-note-app/models"
)

type UserService interface {
	Create(user dto.RegisterUserDto) (models.User, error)
	// Create(user dto.RegisterUserDto) (dto.RegisterUserDto, error)
	// Create(person *model.Person) (*model.Person, error)
}
