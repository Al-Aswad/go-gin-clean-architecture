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
	UpdateNoteByID(ctx *gin.Context)
	DeteleNoteByID(ctx *gin.Context)
	FindNoteByID(ctx *gin.Context)
	All(ctx *gin.Context)
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

	noteAddDto.UserID = convertUserID

	createNote, err := n.noteService.Create(noteAddDto)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to create note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, createNote)
	ctx.JSON(http.StatusCreated, res)

}

func (n *noteController) UpdateNoteByID(ctx *gin.Context) {
	noteGetByIDDto := dto.NoteUpdateByIDDTO{}

	errDto := ctx.ShouldBind(&noteGetByIDDto)
	if errDto != nil {
		res := helpers.BuildErrorResponse("failed to bind request", errDto.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.BuildErrorResponse("failed to convert id", err.Error(), nil)
		return
	}

	updateNote, err := n.noteService.UpdateNoteByID(id, noteGetByIDDto)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to update note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, updateNote)
	ctx.JSON(http.StatusOK, res)

}

func (c *noteController) FindNoteByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.BuildErrorResponse("failed to convert id", err.Error(), nil)
		return
	}

	note, err := c.noteService.FindNoteByID(id)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to find note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, note)
	ctx.JSON(http.StatusOK, res)

}

func (c *noteController) DeteleNoteByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.BuildErrorResponse("failed to convert id", err.Error(), nil)
		return
	}

	_, errFind := c.noteService.FindNoteByID(id)
	if errFind != nil {
		res := helpers.BuildErrorResponse("failed to find note", errFind.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	deleteNote := c.noteService.DeteleNoteByID(id)

	res := helpers.BuildResponse(true, "success", nil, deleteNote)
	ctx.JSON(http.StatusOK, res)
}

func (c *noteController) All(ctx *gin.Context) {
	notes, err := c.noteService.All()
	if err != nil {
		res := helpers.BuildErrorResponse("failed to find note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, notes)
	ctx.JSON(http.StatusOK, res)

}

func (c *noteController) getUserIDbyToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		log.Println(err)
	}

	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])

}
