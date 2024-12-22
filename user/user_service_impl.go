package user

import "gorm.io/gorm"

type ServiceImpl struct {
	userRepository Repository
	dbConnection   *gorm.DB
}

func NewService(userRepository Repository, dbConnection *gorm.DB) *ServiceImpl {
	return &ServiceImpl{
		userRepository: userRepository,
		dbConnection:   dbConnection,
	}
}
