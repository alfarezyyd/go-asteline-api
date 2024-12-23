package user

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user/dto"
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
