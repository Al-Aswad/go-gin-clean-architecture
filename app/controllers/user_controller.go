package controllers

import (
	"go-gin-clean-architecture/app/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func CreateUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
