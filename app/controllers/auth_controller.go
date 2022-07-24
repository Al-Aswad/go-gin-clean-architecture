package controllers

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/models"
	"go-gin-clean-architecture/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authSerive services.AuthSevice
	jwtService services.JWTservice
}

func CreateAuthController(authService services.AuthSevice, jwtService services.JWTservice) AuthController {
	return &authController{
		authSerive: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
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
		ctx.SetCookie("token", generateToken, 3600, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Invalid Credential", "", helpers.EmptyResponse{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}
