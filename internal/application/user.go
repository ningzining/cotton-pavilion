package application

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginDTO struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRet struct {
	Token string `json:"token"`
}

type IUserApplication interface {
	Register(dto RegisterDTO) error
	Login(dto LoginDTO) (*LoginRet, error)
}
