package model

type Notes struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
