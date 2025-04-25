package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"test":    "Pieash",
		})
	})

	router.GET("/me/:id/:newId", func(c *gin.Context) {
		id := c.Param("id")
		newId := c.Param("newId")

		c.JSON(http.StatusOK, gin.H{
			"id":    id,
			"newId": newId,
		})
	})

	router.POST("/me", func(c *gin.Context) {
		type MeRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var meRequest MeRequest

		if err := c.BindJSON(&meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	router.Run(":5000")
}
