package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/mongodb"
	"github.com/lgukasyan/SkyCrypt/pkg/config"
)

func main() {
	// Load environmental variables
	config.LoadEnv()

	// Database connection
	var db mongodb.IDatabaseInterfaceProtocol = mongodb.NewMongoClient(os.Getenv("MONGODB_URI"), os.Getenv("DB_NAME"))
	db.Connect()

	// Database close
	defer db.Disconnect()

	// Router API
	var router *gin.Engine = gin.Default()

	// Endpoints
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Server start
	router.Run(":" + os.Getenv("PORT"))
}
