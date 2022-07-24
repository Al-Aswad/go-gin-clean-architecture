package dto

type LoginDto struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email, min:1"`
	Password string `json:"password" form:"password" binding:"required" validate:"required,min=8"`
}

type LoginDtoOk struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email, min:1"`
	Password string `json:"password" form:"password" binding:"required" validate:"required,min=8"`
	Token    string
}
