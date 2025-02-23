package routes

import (
	"fmt"
	"net/http"

	"apilibrary.com/models"
	"apilibrary.com/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the requested data "})
		return
	}

	err = user.UserSave()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not sign up user"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data"})
		return
	}

	err = user.ValidateCredentials()
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not loging. check credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not loging. check credentials"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login was successful", "token": token})
}
