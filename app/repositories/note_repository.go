package repositories

import (
	"go-gin-note-app/app/dto"
	"go-gin-note-app/app/models"
)

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	UpdateNoteByID(id int, note models.Note) (models.Note, error)
	ArchiveNote(id int, note dto.NoteArhiveDTO) (models.Note, error)
	DeteleNoteByID(id int) bool
	FindNoteByID(id int) (models.Note, error)
	All(userID int) ([]models.Note, error)
}
