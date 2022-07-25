package dto

type NoteUpdateByIDDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"desc"`
	IsArchive   bool   `json:"is_archive"`
}
