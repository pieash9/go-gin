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

func (n *NotesServices) GetNotes() []Note {
	data := []Note{
		{
			Id:   1,
			Name: "Note 1",
		},
		{
			Id:   2,
			Name: "Note 2",
		},
	}
	return data
}

func (n *NotesServices) CreateNote() string {
	err := n.db.Create(&model.Notes{
		Id:     1,
		Title:  "TEST 1",
		Status: true,
	})

	if err != nil {
		fmt.Println(err)
	}

	return "Notes created."
}
