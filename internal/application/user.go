package application

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IUserApplication interface {
	Register(dto RegisterDTO) error
}
