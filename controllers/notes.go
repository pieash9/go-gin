package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/services"
)

type NotesController struct {
	NotesService *services.NotesServices
}

func (n *NotesController) InitNotesController(router *gin.Engine, notesService services.NotesServices) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNote())
	n.NotesService = &notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		order := c.Query("order")

		if order != "ASC" && order != "DESC" {
			order = "ASC" // default
		}
		actualStatus, err := strconv.ParseBool(status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		notes, err := n.NotesService.GetNotes(actualStatus, order)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(200, gin.H{"notes": notes})
	}
}

func (n *NotesController) CreateNote() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.NotesService.CreateNote(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})
	}
}
