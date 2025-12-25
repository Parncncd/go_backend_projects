package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	// todo
	// server.POST("/login", createEvent)
	// server.POST("/events/:id/register", createEvent)
	// server.DELETE("/events/:id/register", createEvent)
}
