package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/controllers"
	internal "github.com/pieash9/go-gin/internal/database"
	"github.com/pieash9/go-gin/services"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	db := internal.InitDb()
	if db == nil {
		return
	}

	notesService := &services.NotesServices{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{
		NotesService: notesService, // 🛠 Inject the service properly
	}
	notesController.InitNotesController(router)

	router.Run(":5000")
}
