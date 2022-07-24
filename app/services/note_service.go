package services

import (
	"go-gin-clean-architecture/app/dto"
	"go-gin-clean-architecture/app/models"
)

type NoteService interface {
	Create(note dto.NoteAddDto) (models.Note, error)
}
