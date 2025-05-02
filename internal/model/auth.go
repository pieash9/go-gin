package model

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
