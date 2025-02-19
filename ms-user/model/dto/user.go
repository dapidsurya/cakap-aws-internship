package dto

type UserDto struct {
	ID       int          `json:"id"`
	Username string       `json:"username"`
	Fullname string       `json:"fullname"`
	Email    string       `json:"email"`
	Products []ProductDto `json:"products"`
}
