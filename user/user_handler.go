package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-asteline-api/config"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"go-asteline-api/user/dto"
	"io"
	"net/http"
)

type Handler struct {
	UserService Service
	viperConfig *viper.Viper
}

func NewHandler(userService Service, viperConfig *viper.Viper) *Handler {
	return &Handler{
		UserService: userService,
		viperConfig: viperConfig,
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
	googleLoginUrl := config.IdentityProviderHolder.GoogleLoginConfig.AuthCodeURL("randomstate")
	ginContext.Redirect(http.StatusSeeOther, googleLoginUrl)
	ginContext.JSON(http.StatusOK, googleLoginUrl)
}

func (userHandler *Handler) ProviderCallback(ginContext *gin.Context) {
	state := ginContext.Query("state")
	if state != "randomstate" {
		exception.ThrowClientError(&exception.ClientError{
			StatusCode: http.StatusUnauthorized,
			Message:    "State does not match",
		})
	}

	code := ginContext.Query("code")

	googlecon := config.GoogleConfig(userHandler.viperConfig)

	token, err := googlecon.Exchange(context.Background(), code)
	helper.CheckErrorOperation(err, &exception.ClientError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Token does not match",
	})
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	helper.CheckErrorOperation(err, &exception.ClientError{
		StatusCode: http.StatusUnauthorized,
		Message:    "User data fetch failed",
	})

	userData, err := io.ReadAll(resp.Body)
	helper.CheckErrorOperation(err, &exception.ClientError{
		StatusCode: http.StatusUnauthorized,
		Message:    "JSON Parsing failed",
	})

	ginContext.JSON(http.StatusOK, string(userData))
}

func (userHandler *Handler) LoginProviderSuccess(ginContext *gin.Context) {}
