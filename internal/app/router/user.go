package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/SkyCrypt/internal/app/controller"
	"github.com/lgukasyan/SkyCrypt/internal/app/service"
	"github.com/lgukasyan/SkyCrypt/internal/middleware"
	repository "github.com/lgukasyan/SkyCrypt/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(router *gin.Engine, db *mongo.Database) {
	var coll *mongo.Collection = db.Collection("users")
	var userRepository repository.IUserRepositoryInterface = repository.NewUserRepository(coll)
	var userService service.IUserServiceInterface = service.NewUserService(userRepository)
	var userController *controller.UserController = controller.NewUserController(userService)

	// Router group
	user := router.Group("/user")

	// Routes
	user.POST("/sign-up", userController.SignUp)
	user.POST("/sign-in", userController.SignIn)

	// Auth middleware
	user.Use(middleware.AuthMiddleware())

	// Auth required
	user.GET("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "auth",
		})
	})
}
