package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/model"
	"go-asteline-api/user/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type ServiceImpl struct {
	userRepository  Repository
	dbConnection    *gorm.DB
	structValidator *validator.Validate
	viperConfig     *viper.Viper
}

func NewService(userRepository Repository, dbConnection *gorm.DB, structValidator *validator.Validate, viperConfig *viper.Viper) *ServiceImpl {
	return &ServiceImpl{
		userRepository:  userRepository,
		dbConnection:    dbConnection,
		structValidator: structValidator,
		viperConfig:     viperConfig,
	}
}

func (userService *ServiceImpl) HandleSave(ginContext *gin.Context, userRegisterDto *dto.UserRegisterDto) bool {
	err := userService.structValidator.Struct(userRegisterDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	dbTransaction := userService.dbConnection.Begin()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterDto.Password), 14)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	userRegisterDto.Password = string(hashedPassword)
	userModel, _ := mapper.MapUserRegisterDtoIntoUserModel(*userRegisterDto)
	defer dbTransaction.Rollback()
	dbTransaction.Create(userModel)
	helper.TransactionOperation(dbTransaction, ginContext)
	return true
}

func (userService *ServiceImpl) HandleLogin(ginContext *gin.Context, userLoginDto *dto.UserLoginDto) bool {
	err := userService.structValidator.Struct(userLoginDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	dbTransaction := userService.dbConnection.Begin()
	var searchedUsers model.User
	searchResult := dbTransaction.Where("email = ?", userLoginDto.Email).First(&searchedUsers).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(searchResult))
	err = bcrypt.CompareHashAndPassword([]byte(searchedUsers.Password), []byte(userLoginDto.Password))
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrUnauthorized))
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userLoginDto.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})
	fmt.Println(jwtToken, userService.viperConfig.GetString("JWT_SECRET"))
	tokenString, err := jwtToken.SignedString([]byte(userService.viperConfig.GetString("JWT_SECRET")))
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrUnauthorized))
	fmt.Println(tokenString)
	ginContext.JSON(200, gin.H{
		"token": tokenString,
	})
	return true
}
