package routes

import (
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(context *gin.Context) {
	userId := context.GetInt64("userId")
	value, _ := context.Get("event")
	event := value.(*models.Event)

	err := event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered on event"})
}

func cancelRegister(context *gin.Context) {
	userId := context.GetInt64("userId")
	value, _ := context.Get("event")
	event := value.(*models.Event)

	err := event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled register on event"})
}
