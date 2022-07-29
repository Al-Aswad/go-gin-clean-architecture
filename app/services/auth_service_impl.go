package services

import (
	"go-gin-note-app/app/models"
	"go-gin-note-app/app/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	authRepo repositories.UserRepository
}

func CreateAuthService(authRepo repositories.UserRepository) AuthSevice {
	return &AuthServiceImpl{
		authRepo: authRepo,
	}
}

func (service *AuthServiceImpl) VerifyCredential(email string, password string) interface{} {
	res := service.authRepo.VerifyUser(email, password)

	if v, ok := res.(models.User); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))

		if v.Email == email && comparedPassword {
			return res
		}

		return false
	}

	return false
}

func comparedPassword(hashPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
