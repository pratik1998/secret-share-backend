/*
	Description: This file contains the route for the hello world api (/api/v1/hello)
	Authored By: Pratik Kayastha
	Copyright 2020 Pratik Kayastha
*/

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGreetings() string {
	return "Hello World"
}

func InitHelloRoute(route *gin.Engine) {

	groupRoute := route.Group("/api/v1")
	groupRoute.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": getGreetings(),
		})
	})
}
