package routes

import (
	"realChakrawarti/rest-event/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Setup handler, GET, POST, PUT, PATCH, DELETE
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authRouteGroup := server.Group("/")
	authRouteGroup.Use(middlewares.Authenticate)

	// server.POST("/events", middlewares.Authenticate, createEvent)
	authRouteGroup.POST("/events", createEvent)
	authRouteGroup.PUT("/events/:id", updateEvent)
	authRouteGroup.DELETE("/events/:id", deleteEvent)
	authRouteGroup.POST("/events/:id/register", registerForEvent)
	authRouteGroup.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
