package user

import (
	"go-asteline-api/repository/user"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	UserRepository user.Repository
	dbConnection   *gorm.DB
}

func NewService(userRepository user.Repository, dbConnection *gorm.DB) *ServiceImpl {
	return &ServiceImpl{
		UserRepository: userRepository,
		dbConnection:   dbConnection,
	}
}
