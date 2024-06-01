package main

import (
	"net/http"
	"realChakrawarti/rest-event/db"
	"realChakrawarti/rest-event/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDb()

	server := gin.Default()

	// Setup handler, GET, POST, PUT, PATCH, DELETE
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":42069") // localhost:42069
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find events. Please try again after sometime!",
		})
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data",
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the event. Please try again after sometime",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"data":    event,
	})
}
