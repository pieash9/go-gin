package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/controllers"
)

func main() {
	router := gin.Default()

	notestController := &controllers.NotestsController{}
	notestController.InitNotesController(router)

	router.Run(":5000")
}
