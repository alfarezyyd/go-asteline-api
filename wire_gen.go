// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-asteline-api/config"
	"go-asteline-api/user"
	"gorm.io/gorm"
)

// Injectors from injector.go:

func InitializeUserController(gormConnection *gorm.DB, validatorInstance *validator.Validate) user.Controller {
	repositoryImpl := user.NewRepository()
	serviceImpl := user.NewService(repositoryImpl, gormConnection, validatorInstance)
	handler := user.NewHandler(serviceImpl)
	return handler
}

// injector.go:

var userFeatureSet = wire.NewSet(config.NewDatabaseConnection, user.NewRepository, wire.Bind(new(user.Repository), new(*user.RepositoryImpl)), user.NewService, wire.Bind(new(user.Service), new(*user.ServiceImpl)), user.NewHandler, wire.Bind(new(user.Controller), new(*user.Handler)))
