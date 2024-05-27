package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Setup handler, GET, POST, PUT, PATCH, DELETE
	server.GET("/events", getEvents)

	server.Run(":42069") // localhost:42069
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello!",
	})
}