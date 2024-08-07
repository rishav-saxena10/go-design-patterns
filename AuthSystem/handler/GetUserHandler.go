package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": nil,
	})
}
