package services

import (
	"gin-note-app/dto"
	"gin-note-app/models"
)

type NoteService interface {
	Create(note dto.NoteAddDto) (models.Note, error)
	UpdateNoteByID(id int, note dto.NoteUpdateByIDDTO) (models.Note, error)
	UpdateArchive(id int, note dto.NoteArhiveDTO) (models.Note, error)
	DeteleNoteByID(id int) bool
	FindNoteByID(id int) (models.Note, error)
	All(userID int) ([]models.Note, error)
}
