package dto

type RegisterUserDto struct {
	Username string `json:"username" form:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type ResponRegistrasi struct {
	Username string `json:"username" form:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=100"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}
