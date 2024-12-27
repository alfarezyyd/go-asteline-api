package user

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/exception"
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
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	userHandler.UserService.HandleSave(ginContext, &userRegisterDto)
}

func (userHandler *Handler) Login(ginContext *gin.Context) {
	var userLoginDto dto.UserLoginDto
	err := ginContext.ShouldBindJSON(&userLoginDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	userHandler.UserService.HandleLogin(ginContext, &userLoginDto)
}
