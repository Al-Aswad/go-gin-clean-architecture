package repositories

import (
	"gin-note-app/dto"
	"gin-note-app/models"

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

func (n *NoteRepositoryImpl) UpdateNoteByID(id int, note models.Note) (models.Note, error) {
	noteUpdate := models.Note{}

	// Update attributes with `struct`, will only update non-zero fields
	err := n.db.Model(&noteUpdate).Where("id = ?", id).Updates(&note).Error
	if err != nil {
		return models.Note{}, err
	}

	return noteUpdate, nil
}
func (n *NoteRepositoryImpl) ArchiveNote(id int, note dto.NoteArhiveDTO) (models.Note, error) {
	noteUpdate := models.Note{}

	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	errFind := n.db.First(&noteUpdate, "id = ?", id).Error
	if errFind != nil {
		return models.Note{}, errFind
	}
	// Update attributes with `struct`, will only update non-zero fields
	err := n.db.Debug().Model(&noteUpdate).Where("id = ?", id).Update("is_archive", note.IsArchive).Error
	if err != nil {
		return models.Note{}, err
	}

	return noteUpdate, nil
}

func (n *NoteRepositoryImpl) FindNoteByID(id int) (models.Note, error) {
	note := models.Note{}
	err := n.db.Where("id = ?", id).First(&note).Error
	if err != nil {
		return models.Note{}, err
	}
	return note, nil
}

func (n *NoteRepositoryImpl) DeteleNoteByID(id int) bool {
	return n.db.Delete(&models.Note{}, "id = ?", id).RowsAffected > 0
}

func (n *NoteRepositoryImpl) All(userID int) ([]models.Note, error) {
	var notes []models.Note

	err := n.db.Where("user_id = ?", userID).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}
