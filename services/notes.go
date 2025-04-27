package services

import (
	"fmt"

	"github.com/pieash9/go-gin/internal/model"
	"gorm.io/gorm"
)

type NotesServices struct {
	db *gorm.DB
}

func (n *NotesServices) InitService(database *gorm.DB) {
	n.db = database

	n.db.AutoMigrate(&model.Notes{})
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NotesServices) GetNotes(status bool, order string) ([]*model.Notes, error) {
	var notes []*model.Notes

	if err := n.db.Where("status = ?", status).Order("id " + order).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *NotesServices) CreateNote(title string, status bool) (*model.Notes, error) {
	note := &model.Notes{
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return note, nil
}
