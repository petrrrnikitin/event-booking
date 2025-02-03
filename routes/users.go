package routes

import (
	"event-booking/models"
	"event-booking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(
			http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Could not parse request data"})
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

func login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Could not parse request data"})
	}

	err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwt, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "User login successfully!", "accessToken": jwt})

}
