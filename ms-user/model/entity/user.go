package entity

type User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
