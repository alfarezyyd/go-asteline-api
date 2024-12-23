package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/user/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	userRepository  Repository
	dbConnection    *gorm.DB
	structValidator *validator.Validate
}

func NewService(userRepository Repository, dbConnection *gorm.DB, structValidator *validator.Validate) *ServiceImpl {
	return &ServiceImpl{
		userRepository:  userRepository,
		dbConnection:    dbConnection,
		structValidator: structValidator,
	}
}

func (userService *ServiceImpl) HandleSave(ginContext *gin.Context, userRegisterDto *dto.UserRegisterDto) bool {
	err := userService.structValidator.Struct(userRegisterDto)
	if errorResult := helper.CheckErrorOperation(err, func() {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}); errorResult {
		return false
	}
	dbTransaction := userService.dbConnection.Begin()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterDto.Password), 14)
	if errorResult := helper.CheckErrorOperation(err, func() {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}); errorResult {
		return false
	}
	userRegisterDto.Password = string(hashedPassword)
	userModel, _ := mapper.MapUserRegisterDtoIntoUserModel(*userRegisterDto)
	defer dbTransaction.Rollback()
	dbTransaction.Create(userModel)
	helper.TransactionOperation(dbTransaction)
	return true
}
