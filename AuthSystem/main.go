package main

import (
	"learning/go/AuthSystem/handler"
	"learning/go/AuthSystem/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Public Routes
	router.POST("/register", handler.RegisterHandler)
	router.POST("/login", handler.LoginHandler)
	router.GET("/test", testHandler)

	// Protected Routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/user", handler.GetUserHandler)
		// Other protected routes
	}

	router.Run(":9080")
}

func testHandler(c *gin.Context) {
	log.Println("Hello Server")
	c.JSON(http.StatusOK, gin.H{
		"test": "OK",
	})
}
