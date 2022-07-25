package repositories

import (
	"go-gin-clean-architecture/app/models"
)

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	UpdateNoteByID(id int, note models.Note) (models.Note, error)
	DeteleNoteByID(id int) bool
	FindNoteByID(id int) (models.Note, error)
	All() ([]models.Note, error)
}
