package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/services"
)

type NotestsController struct {
	notesService services.NotesServices
}

func (n *NotestsController) InitNotesController(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())

	notes.POST("/", n.CreateNote())
}

func (n *NotestsController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": n.notesService.GetNotes()})
	}
}
func (n *NotestsController) CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": n.notesService.CreateNote()})
	}
}
