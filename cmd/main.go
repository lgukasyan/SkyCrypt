package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	app "github.com/lgukasyan/SkyCrypt/internal/app/router"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/auth"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/mongodb"
	"github.com/lgukasyan/SkyCrypt/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Load environmental variables
	config.LoadEnv()

	// Connect & disconnect database
	var client mongodb.IDatabaseInterfaceProtocol = mongodb.NewMongoClient(os.Getenv("MONGODB_URI"), os.Getenv("DB_NAME"))
	defer client.Disconnect()

	// Set JWT Key
	auth.SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

	// Get Database
	var db *mongo.Database = client.GetDatabase()

	// Router API
	var router *gin.Engine = gin.Default()

	// Setup routes
	app.UserRouter(router, db)

	// Endpoints
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Server start
	router.Run(":" + os.Getenv("PORT"))
}
