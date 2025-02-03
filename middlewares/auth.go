package middlewares

import (
	"event-booking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.GetUserIdFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
