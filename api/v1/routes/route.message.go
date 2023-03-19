package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var messages map[string]string

type Message struct {
	Message string `json:"message"`
}

func getMessage(uuid string) string {
	return messages[uuid]
}

func storeMessge(msg Message) string {
	uuid := uuid.New().String()
	messages[uuid] = msg.Message
	return uuid
}

func InitMessagesRoute(route *gin.Engine) {
	messages = make(map[string]string)
	groupRoute := route.Group("/api/v1")

	// TODO: Add proper authentication and authorization to check if the user is authorized to access the message

	// Get message by uuid
	groupRoute.GET("/message/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		c.JSON(http.StatusOK, gin.H{
			"message": getMessage(uuid),
		})
	})

	// Store message and return uuid
	groupRoute.POST("/message", func(c *gin.Context) {
		var msg Message
		c.BindJSON(&msg)
		uuid := storeMessge(msg)
		c.JSON(http.StatusOK, gin.H{
			"uuid": uuid,
		})
	})
}
