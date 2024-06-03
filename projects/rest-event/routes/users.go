package routes

import (
	"log"
	"net/http"
	"realChakrawarti/rest-event/models"
	"realChakrawarti/rest-event/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user. Please try again after sometime",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data",
		})
		log.Fatal(err)
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful.",
		"token":   token,
	})
}
