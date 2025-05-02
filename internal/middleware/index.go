package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/internal/utils"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")
	fmt.Println(headers)

	if headers == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	token := strings.Split(headers, " ")
	if len(token) != 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token not provided. Unauthorized"})
		return
	}

	data, err := utils.TokenCheck(token[1])
	fmt.Println(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Token not matched. Unauthorized"})
		return
	}

	c.Next()
}
