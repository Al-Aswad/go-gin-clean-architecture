package dto

type NoteAddDto struct {
	Title       string `json:"title" form:"title" binding:"required,min=3,max=100"`
	Body        string `json:"body" form:"body" binding:"required,min=3"`
	Description string `json:"desc" form:"description" binding:"required,min=3"`
}
