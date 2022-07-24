package services

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/models"
	"go-gin-clean-architecture/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type NoteServiceImpl struct {
	noteRepo repositories.NoteRepository
}

func CreateNoteService(noteRepo repositories.NoteRepository) NoteService {
	return &NoteServiceImpl{
		noteRepo: noteRepo,
	}

}

func (usImpl *NoteServiceImpl) Create(note dto.NoteAddDto) (models.Note, error) {
	noteCreate := models.Note{}
	err := smapping.FillStruct(&noteCreate, smapping.MapFields(&note))

	if err != nil {
		log.Println("[NoteServiceImpl.Create] error fill struct", err)
		return noteCreate, err
	}

	log.Println("[NoteServiceImpl.Create] Body", note)
	log.Println("[NoteServiceImpl.Create] noteCreate", noteCreate)

	noteCreate, err = usImpl.noteRepo.Create(noteCreate)
	if err != nil {
		log.Println("[NoteServiceImpl.Create] error execute query", err)
		return models.Note{}, err
	}
	return noteCreate, nil
}
