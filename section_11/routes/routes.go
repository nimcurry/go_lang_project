package routes

import (
	"apilibrary.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // /events/1 /events/5

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvents)
	authenticated.DELETE("/events/:id", deleteEvents)
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/register", cancelRegisterationForEvents)

	server.POST("/users", signup)
	server.POST("/login", login)
}
