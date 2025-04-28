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

func (n *NotesServices) GetNote(id int) (*model.Notes, error) {
	var note *model.Notes

	if err := n.db.Where("id = ?", id).Find(&note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NotesServices) GetAllNotes(order string) ([]*model.Notes, error) {
	var notes []*model.Notes

	if err := n.db.Order("id " + order).Find(&notes).Error; err != nil {
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

func (n *NotesServices) UpdateNote(title string, status bool, id int) (*model.Notes, error) {

	var note *model.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	if err := n.db.Save(&note).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return note, nil
}

func (n *NotesServices) DeleteNote(id int) error {

	var note *model.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
