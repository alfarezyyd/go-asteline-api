package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	databaseInstance *gorm.DB
}

func NewDatabaseConnection() *DatabaseConnection {
	return &DatabaseConnection{}
}

func (dbConn *DatabaseConnection) GetDatabaseConnection() *gorm.DB {
	if dbConn.databaseInstance == nil {
		sqlDialect := mysql.Open("root:@tcp(127.0.0.1:3306)/go-asteline-api?charset=utf8mb4&parseTime=True&loc=Local")
		gormOpen, err := gorm.Open(sqlDialect, &gorm.Config{})
		dbConn.databaseInstance = gormOpen
		if err != nil {
			panic(err)
		}
	}
	return dbConn.databaseInstance
}
