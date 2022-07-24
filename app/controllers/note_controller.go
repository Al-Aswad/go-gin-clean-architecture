package controllers

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/helpers"
	"go-gin-clean-architecture/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NoteController interface {
	Create(ctx *gin.Context)
}

type noteController struct {
	// Service
	noteService services.NoteService
}

func CreateNoteController(noteService services.NoteService) NoteController {
	return &noteController{
		noteService: noteService,
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

	createNote, err := n.noteService.Create(noteAddDto)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to create note", err.Error(), nil)

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "success", nil, createNote)
	ctx.JSON(http.StatusCreated, res)

}
