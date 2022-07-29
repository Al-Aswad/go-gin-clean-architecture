package services

import (
	"go-gin-note-app/app/dto"
	"go-gin-note-app/app/models"
	"go-gin-note-app/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func CreateUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}

}

func (u *UserServiceImpl) Create(user dto.RegisterUserDto) (models.User, error) {
	userCreate := models.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))

	if err != nil {
		log.Println("[UserServiceImpl.Create] error fill struct", err)
		return userCreate, err
	}

	log.Println("[UserServiceImpl.Create] userCreate", userCreate)

	userCreate, err = u.userRepo.Create(userCreate)
	if err != nil {
		log.Println("[UserServiceImpl.Create] error execute query", err)
		return models.User{}, err
	}
	return userCreate, nil
}
