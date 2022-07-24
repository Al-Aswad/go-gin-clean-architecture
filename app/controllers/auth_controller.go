package controllers

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/models"
	"go-gin-clean-architecture/app/services"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authSerive  services.AuthSevice
	userService services.UserService
	jwtService  services.JWTservice
}

func CreateAuthController(authService services.AuthSevice, userService services.UserService, jwtService services.JWTservice) AuthController {
	return &authController{
		authSerive:  authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	var loginDTO dto.LoginDto

	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to bind request", errDTO.Error(), helpers.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authResult := c.authSerive.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		response := helpers.BuildResponse(true, "OK!", nil, v)
		baseUrl := os.Getenv("DOMAIN")
		log.Println("Domain Cookie ", baseUrl)
		ctx.SetCookie("token", generateToken, 3600, "/", baseUrl, false, true)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Invalid Credential", "", helpers.EmptyResponse{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}

func (c *authController) Register(ctx *gin.Context) {
	var dtoUserCreate dto.RegisterUserDto

	errDto := ctx.ShouldBind(&dtoUserCreate)
	if errDto != nil {
		res := helpers.BuildErrorResponse("failed to bind request", errDto.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createUser, err := c.userService.Create(dtoUserCreate)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to create user", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, createUser)
	ctx.JSON(http.StatusCreated, res)

}
