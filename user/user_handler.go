package user

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/helper"
	"go-asteline-api/user/dto"
	"net/http"
)

type Handler struct {
	UserService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (userHandler *Handler) Register(ginContext *gin.Context) {
	var userRegisterDto dto.UserRegisterDto
	err := ginContext.ShouldBindJSON(&userRegisterDto)
	userHandler.UserService.HandleSave(ginContext, &userRegisterDto)

	if err != nil {
		panic(err)
		return
	}
}

func (userHandler *Handler) Login(ginContext *gin.Context) {
	var userLoginDto dto.UserLoginDto
	err := ginContext.ShouldBindJSON(&userLoginDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest, err.Error()) {
		return
	}
	userHandler.UserService.HandleLogin(ginContext, &userLoginDto)
}
