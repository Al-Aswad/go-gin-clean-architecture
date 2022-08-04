package dto

type NoteUpdateByIDDTO struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"desc"`
	IsArchive   bool   `json:"is_archive"`
}

type NoteArhiveDTO struct {
	IsArchive bool `json:"is_archive"`
}
