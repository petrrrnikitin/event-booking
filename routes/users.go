package routes

import (
	"event-booking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Could not save user",
		})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
