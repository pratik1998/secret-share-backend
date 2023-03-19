package main

import (
	route "github.com/pratik1998/secret-share-backend/api/v1/routes"
	util "github.com/pratik1998/secret-share-backend/utils"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		log.Info("Starting server in development mode")
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		log.Info("Starting server in test mode")
		gin.SetMode(gin.TestMode)
	} else {
		log.Info("Starting server in production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	route.InitHelloRoute(router)
	route.InitMessagesRoute(router)

	router.Run(":8080")
}
