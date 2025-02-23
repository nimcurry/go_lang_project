package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"apilibrary.com/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	//context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find any events"})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {

	/* token := context.Request.Header.Get("Authorization")
	fmt.Println("token: ", token)

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized token"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized token"})
		return
	}
	*/
	userId := context.GetInt64("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	event.UserId = userId
	//event.UserId = 1

	fmt.Println("error is:", err)

	err = event.Save()

	fmt.Println("error is:", err)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
	//event.Save()
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse eventId"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find the eventId"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "not able to update the event"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()

	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "not able to update the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated sucessfully"})

}

func deleteEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "not able to find the id.."})
		return
	}
	userId := context.GetInt64("userId")

	deleteEvent, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event"})
		return
	}
	if deleteEvent.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}
	/* 	var deleteEvent models.Event
	   	err = context.ShouldBindJSON(&deleteEvent)
	   	if err != nil {
	   		context.JSON(http.StatusBadRequest, gin.H{"message": "not able to update the event"})
	   		return
	   	}
	   	deleteEvent.ID = eventId */

	err = deleteEvent.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})

}
