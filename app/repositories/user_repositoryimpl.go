package repositories

import (
	"fmt"
	"go-gin-note-app/app/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func CreateUserRepo(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create(user models.User) (models.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))

	err := u.db.Save(&user).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
		return models.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) VerifyUser(email string, password string) interface{} {
	var user models.User
	res := u.db.Where("email= ?", email).Take(&user)

	if res.Error != nil {
		return nil
	}

	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}

	return string(hash)
}
