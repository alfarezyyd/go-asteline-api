package user

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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

func (userHandler *Handler) LoginWithProvider(ginContext *gin.Context) {
	provider := ginContext.Param("provider")
	queryRequest := ginContext.Request.URL.Query()
	queryRequest.Add("provider", provider)
	ginContext.Request.URL.RawQuery = queryRequest.Encode()
	gothic.BeginAuthHandler(ginContext.Writer, ginContext.Request)
}

func (userHandler *Handler) ProviderCallback(ginContext *gin.Context) {
	loginProvider := ginContext.Param("provider")
	queryRequest := ginContext.Request.URL.Query()
	queryRequest.Add("provider", loginProvider)
	ginContext.Request.URL.RawQuery = queryRequest.Encode()

	_, err := gothic.CompleteUserAuth(ginContext.Writer, ginContext.Request)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, "Internal server error"))
	ginContext.Redirect(http.StatusTemporaryRedirect, "/success")
}

func (userHandler *Handler) LoginProviderSuccess(ginContext *gin.Context) {}
