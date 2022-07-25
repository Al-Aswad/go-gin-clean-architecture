package services

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/models"
)

type NoteService interface {
	Create(note dto.NoteAddDto) (models.Note, error)
	UpdateNoteByID(id int, note dto.NoteUpdateByIDDTO) (models.Note, error)
	DeteleNoteByID(id int) bool
	FindNoteByID(id int) (models.Note, error)
	All(userID int) ([]models.Note, error)
}
