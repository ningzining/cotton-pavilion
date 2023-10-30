package entity

type User struct {
	BaseModel
	Username string
	Mobile   string
	Email    string
	Password string
}
