package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/internal/model"
	"github.com/pieash9/go-gin/services"
)

type NotesController struct {
	NotesService *services.NotesServices
}

func (n *NotesController) InitController(notesService services.NotesServices) *NotesController {

	// return &NotesController{
	// 	NotesService: &notesService,
	// }
	n.NotesService = &notesService
	return n
}

func (n *NotesController) InitRoutes(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetNote())
	notes.POST("/", n.CreateNote())
	notes.PUT("/", n.UpdateNote())
	notes.DELETE("/:id", n.DeleteNote())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		order := c.Query("order")

		if order != "asc" && order != "desc" {
			order = "asc" // default
		}

		var notes []*model.Notes
		var err error

		if status == "" {
			// No status filter
			notes, err = n.NotesService.GetAllNotes(order)
		} else {
			actualStatus, errParse := strconv.ParseBool(status)
			if errParse != nil {
				c.JSON(400, gin.H{
					"message": errParse.Error(),
				})
				return
			}
			notes, err = n.NotesService.GetNotes(actualStatus, order)
		}

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"notes": notes})
	}
}

func (n *NotesController) GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		var note *model.Notes

		note, err = n.NotesService.GetNote(id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"notes": note})
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

func (n *NotesController) UpdateNote() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.NotesService.UpdateNote(noteBody.Title, noteBody.Status, noteBody.Id)
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

func (n *NotesController) DeleteNote() gin.HandlerFunc {

	return func(c *gin.Context) {
		idParam := c.Param("id") // or c.Query("id") depending where it comes from
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid ID"})
			return
		}

		err = n.NotesService.DeleteNote(id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Note Deleted!",
		})
	}
}
