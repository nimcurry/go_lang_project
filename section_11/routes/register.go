package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"apilibrary.com/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 0, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	fmt.Println("user and event", userId, event)
	err = event.Register(userId)
	fmt.Println("error: ", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not insert into event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registered"})

}

func cancelRegisterationForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 0, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "some issue occured"})
		return
	}
	var event models.Event

	event.ID = eventId

	err = event.Cancel(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel the registrations"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registration cnaceled"})
}
