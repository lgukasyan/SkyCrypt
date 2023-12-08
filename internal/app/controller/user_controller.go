package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	domain "github.com/lgukasyan/SkyCrypt/domain/user"
	"github.com/lgukasyan/SkyCrypt/internal/app/service"
	"github.com/lgukasyan/SkyCrypt/internal/infrastructure/auth"
	"github.com/lgukasyan/SkyCrypt/internal/response"
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
		response.Write(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	if err := uc.userService.InsertUser(ctx, &user); err != nil {
		response.Write(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	response.Write(ctx, http.StatusCreated, nil, "user created successfully.")
}

func (uc *UserController) SignIn(ctx *gin.Context) {
	var user domain.UserSignIn

	if err := jsoniter.ConfigFastest.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		response.Write(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	doc, err := uc.userService.FindAndValidate(ctx, &user)
	if err != nil {
		response.Write(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	var token domain.Token
	token.AccessToken, err = auth.EncodeJWT(doc.Id, 15)
	if err != nil {
		response.Write(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	token.RefreshToken, err = auth.EncodeJWT(doc.Id, 1440)
	if err != nil {
		response.Write(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.Write(ctx, http.StatusCreated, token, "user logged in successfully.")
}
