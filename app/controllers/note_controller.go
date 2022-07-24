package controllers

import (
	"fmt"
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type NoteController interface {
	Create(ctx *gin.Context)
}

type noteController struct {
	// Service
	noteService services.NoteService
	jwtService  services.JWTservice
}

func CreateNoteController(noteService services.NoteService, jwtSevice services.JWTservice) NoteController {
	return &noteController{
		noteService: noteService,
		jwtService:  jwtSevice,
	}
}

func (n *noteController) Create(ctx *gin.Context) {
	noteAddDto := dto.NoteAddDto{}

	errDto := ctx.ShouldBind(&noteAddDto)
	if errDto != nil {
		res := helpers.BuildErrorResponse("failed to bind request", errDto.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// userId, err := helpers.GetUserIdFromContext(ctx)
	createNote, err := n.noteService.Create(noteAddDto)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to create note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	cookieToken, err := ctx.Cookie("token")
	if err != nil {
		res := helpers.BuildErrorResponse("Token Not Found", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userID := n.getUserIDbyToken(cookieToken)
	convertUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to convert user id", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createNote.UserID = convertUserID
	res := helpers.BuildResponse(true, "success", nil, createNote)
	ctx.JSON(http.StatusCreated, res)

}

func (c *noteController) getUserIDbyToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		log.Println(err)
	}

	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])

}
