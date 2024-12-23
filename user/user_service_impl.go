package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return false
	}
	dbTransaction := userService.dbConnection.Begin()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterDto.Password), 14)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return false
	}
	userRegisterDto.Password = string(hashedPassword)
	userModel, _ := mapper.MapUserRegisterDtoIntoUserModel(*userRegisterDto)
	defer dbTransaction.Rollback()
	dbTransaction.Create(userModel)
	helper.TransactionOperation(dbTransaction)
	return true
}

func (userService *ServiceImpl) HandleLogin(ginContext *gin.Context, userLoginDto *dto.UserLoginDto) bool {
	err := userService.structValidator.Struct(userLoginDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return false
	}
	dbTransaction := userService.dbConnection.Begin()
	var searchedUsers model.User
	searchResult := dbTransaction.Where("email = ?", userLoginDto.Email).First(&searchedUsers)
	if helper.CheckErrorOperation(searchResult.Error, ginContext, http.StatusBadRequest) {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(searchedUsers.Password), []byte(userLoginDto.Password))
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return false
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userLoginDto.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})
	fmt.Println(jwtToken, userService.viperConfig.GetString("JWT_SECRET"))
	tokenString, err := jwtToken.SignedString([]byte(userService.viperConfig.GetString("JWT_SECRET")))
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return false
	}
	fmt.Println(tokenString)
	ginContext.JSON(200, gin.H{
		"token": tokenString,
	})
	return true
}
