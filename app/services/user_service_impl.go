package services

import (
	"go-gin-clean-architecture/app/models"
	"go-gin-clean-architecture/app/repositories"
)

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func CreateUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}

}

func (u *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	return u.userRepo.Create(user)
}
