package user

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user/dto"
)

type Service interface {
	HandleSave(ginContext *gin.Context, userRegisterDto *dto.UserRegisterDto) bool
	HandleLogin(ginContext *gin.Context, userLoginDto *dto.UserLoginDto) bool
}
