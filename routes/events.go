package routes

import (
	"event-booking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid id: %d", id)})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect data, some fields are missing", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.CreatorID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func updateEvent(context *gin.Context) {
	value, _ := context.Get("event")
	event := value.(*models.Event)

	userId := context.GetInt64("userId")

	if event.CreatorID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this event"})
		return
	}

	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect data, some fields are missing", "error": err.Error()})
		return
	}

	updatedEvent.ID = event.ID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	value, _ := context.Get("event")
	event := value.(*models.Event)

	userId := context.GetInt64("userId")

	if event.CreatorID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this event"})
		return
	}

	err := event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
