package repositories

import (
	"go-gin-note-app/app/models"
)

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	UpdateNoteByID(id int, note models.Note) (models.Note, error)
	DeteleNoteByID(id int) bool
	FindNoteByID(id int) (models.Note, error)
	All(userID int) ([]models.Note, error)
}
