package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/services"
)

type NotesController struct {
	NotesService *services.NotesServices
}

func (n *NotesController) InitNotesController(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())

	notes.POST("/", n.CreateNote())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": n.NotesService.GetNotes()})
	}
}
func (n *NotesController) CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": n.NotesService.CreateNote()})
	}
}
