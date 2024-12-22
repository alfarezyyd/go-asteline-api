package user

import (
	"gorm.io/gorm"
)

type ServiceImpl struct {
	UserRepository Repository
	dbConnection   *gorm.DB
}

func NewService(userRepository Repository, dbConnection *gorm.DB) Service {
	return &ServiceImpl{
		UserRepository: userRepository,
		dbConnection:   dbConnection,
	}
}
