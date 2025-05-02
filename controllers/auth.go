package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pieash9/go-gin/services"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")

	routes.POST("/login", a.Nope())
	routes.POST("/register", a.Register())
}

func (*AuthController) Nope() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Connected",
		})
		return
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBody
		if err := c.BindJSON(&registerBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": user,
		})
	}
}
