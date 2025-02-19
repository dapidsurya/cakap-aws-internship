package entity

type User struct {
	ID       int    `json:"id" gorm:"column:id_user"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
