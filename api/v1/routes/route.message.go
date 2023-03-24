package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/pratik1998/secret-share-backend/controllers"
)

type MessageBody struct {
	Message string `json:"message"`
}

func getMessage(uuid string) string {
	return controllers.GetMessage(uuid)
}

func storeMessge(msg string) string {
	return controllers.StoreMessage(msg)
}

func InitMessagesRoute(route *gin.Engine) {
	groupRoute := route.Group("/api/v1")

	// TODO: Add proper authentication and authorization to check if the user is authorized to access the message

	// Get message by uuid
	groupRoute.GET("/message/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		message := getMessage(uuid)
		if message == "Invalid UUID" {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	// Store message and return uuid
	groupRoute.POST("/message", func(c *gin.Context) {
		var msg MessageBody
		c.BindJSON(&msg)
		uuid := storeMessge(msg.Message)
		c.JSON(http.StatusOK, gin.H{
			"uuid": uuid,
		})
	})
}
