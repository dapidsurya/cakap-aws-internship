package entity

type Product struct {
	UserID int    `json:"userId" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	Stock  int    `json:"stock" gorm:"column:stock"`
}

func (Product) TableName() string {
	return "product"
}
