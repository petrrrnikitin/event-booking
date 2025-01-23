package main

import (
	"event-booking/db"
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/events", getEvents)
	r.POST("/events", createEvent)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect data, some fields are missing", "error": err.Error()})
		return
	}

	event.ID = len(models.GetAllEvents()) + 1
	event.CreatorID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
