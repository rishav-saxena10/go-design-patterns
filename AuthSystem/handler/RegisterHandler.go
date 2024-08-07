package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterHandler handles the user registration.
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest

	// Bind the JSON request body to the RegisterRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the user
	// user, err := RegisterUser(req.Name, req.Email, req.Password)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
	// 	return
	// }

	// Respond with the created user details (without the password)
	c.JSON(http.StatusCreated, gin.H{
		"id":    "userId",
		"name":  "userName",
		"email": "userEmail",
	})
}
