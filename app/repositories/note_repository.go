package repositories

import "go-gin-clean-architecture/app/models"

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
}
