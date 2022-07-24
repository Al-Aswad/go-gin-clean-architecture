package controllers

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/services"
	"net/http"

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
	var dtoUserCreate dto.RegisterUserDto

	errDto := ctx.ShouldBind(&dtoUserCreate)
	if errDto != nil {
		res := helpers.BuildErrorResponse("failed to bind request", errDto.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createUser, err := u.userService.Create(dtoUserCreate)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to create user", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, createUser)
	ctx.JSON(http.StatusCreated, res)

}
