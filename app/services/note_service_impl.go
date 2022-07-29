package services

import (
	"go-gin-note-app/app/dto"
	"go-gin-note-app/app/models"
	"go-gin-note-app/app/repositories"
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

func (n *NoteServiceImpl) UpdateNoteByID(id int, note dto.NoteUpdateByIDDTO) (models.Note, error) {
	noteUpdate := models.Note{}
	err := smapping.FillStruct(&noteUpdate, smapping.MapFields(&note))
	if err != nil {
		log.Println("[NoteServiceImpl.Create] error fill struct", err)
		return noteUpdate, err
	}

	noteUpdate, err = n.noteRepo.UpdateNoteByID(id, noteUpdate)

	log.Println("[NoteServiceImpl.UpdateNoteByID] Body", noteUpdate)

	if err != nil {
		return models.Note{}, err
	}
	return noteUpdate, nil
}

func (n *NoteServiceImpl) FindNoteByID(id int) (models.Note, error) {
	note, err := n.noteRepo.FindNoteByID(id)
	if err != nil {
		return models.Note{}, err
	}
	return note, nil
}

func (n *NoteServiceImpl) DeteleNoteByID(id int) bool {
	return n.noteRepo.DeteleNoteByID(id)
}

func (n *NoteServiceImpl) All(userID int) ([]models.Note, error) {
	notes, err := n.noteRepo.All(userID)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
