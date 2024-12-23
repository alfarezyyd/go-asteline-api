//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-asteline-api/config"
	"go-asteline-api/user"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	config.NewDatabaseConnection,
	user.NewRepository,
	wire.Bind(new(user.Repository), new(*user.RepositoryImpl)),
	user.NewService,
	wire.Bind(new(user.Service), new(*user.ServiceImpl)),
	user.NewHandler,
	wire.Bind(new(user.Controller), new(*user.Handler)),
)

func InitializeUserController(gormConnection *gorm.DB, validatorInstance *validator.Validate) user.Controller {
	wire.Build(userFeatureSet)
	return nil
}
