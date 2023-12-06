package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	domain "github.com/lgukasyan/SkyCrypt/domain/user"
	"github.com/lgukasyan/SkyCrypt/internal/app/service"
)

type UserController struct {
	userService service.IUserServiceInterface
}

func NewUserController(userService service.IUserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) SignUp(ctx *gin.Context) {
	var user domain.User
	if err := jsoniter.ConfigFastest.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return
	}

	if err := uc.userService.InsertUser(ctx, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		return
	}

	ctx.JSON(200, gin.H{"user_data": user})
}