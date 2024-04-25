package api

import (
	"crabi-tech-test/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(service application.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := service.CreateUser(req.ID, req.Name, req.LastName, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func Login(service application.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if c.Bind(&loginData) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		}
		user, err := service.Login(loginData.Email, loginData.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUser(service application.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		user, err := service.GetUser(email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
