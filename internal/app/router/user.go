package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/SkyCrypt/internal/app/controller"
	"github.com/lgukasyan/SkyCrypt/internal/app/service"
	repository "github.com/lgukasyan/SkyCrypt/repository/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(router *gin.Engine, db *mongo.Database) {
	var coll *mongo.Collection = db.Collection("users")
	var userRepository repository.IUserRepositoryInterface = repository.NewUserRepository(coll)
	var userService service.IUserServiceInterface = service.NewUserService(userRepository)
	var userController *controller.UserController = controller.NewUserController(userService)

	user := router.Group("/user")
	user.POST("/sign-up", userController.SignUp)
	user.POST("/sign-in", userController.SignIn)
}
