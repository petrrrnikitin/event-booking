package routes

import (
	"event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	auth := server.Group("/")
	auth.Use(middlewares.Auth)
	getEventById := auth.Group("/")
	getEventById.Use(middlewares.LoadEventByRequest)

	auth.POST("/events", createEvent)
	getEventById.PUT("/events/:id", updateEvent)
	getEventById.DELETE("events/:id", deleteEvent)
	getEventById.POST("events/:id/register", register)
	getEventById.DELETE("events/:id/register", cancelRegister)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
