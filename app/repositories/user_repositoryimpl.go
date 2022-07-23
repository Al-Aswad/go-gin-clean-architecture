package repositories

import (
	"fmt"
	"go-gin-clean-architecture/app/models"

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

func (u *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	if err != nil {
		fmt.Printf("[PersonRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return user, nil
}
