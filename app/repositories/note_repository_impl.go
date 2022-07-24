package repositories

import (
	"go-gin-clean-architecture/app/models"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	db *gorm.DB
}

func CreateNoteRepository(db *gorm.DB) *NoteRepositoryImpl {
	return &NoteRepositoryImpl{
		db: db,
	}
}

func (n *NoteRepositoryImpl) Create(note models.Note) (models.Note, error) {
	// note.UserID = 1
	err := n.db.Save(&note).Error
	if err != nil {
		return models.Note{}, err
	}
	return note, nil
}
