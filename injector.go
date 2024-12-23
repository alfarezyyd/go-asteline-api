//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-asteline-api/campaign"
	"go-asteline-api/user"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.Repository), new(*user.RepositoryImpl)),
	user.NewService,
	wire.Bind(new(user.Service), new(*user.ServiceImpl)),
	user.NewHandler,
	wire.Bind(new(user.Controller), new(*user.Handler)),
)

var campaignFeatureSet = wire.NewSet(
	campaign.NewRepository,
	wire.Bind(new(campaign.Repository), new(*campaign.RepositoryImpl)),
	campaign.NewService,
	wire.Bind(new(campaign.Service), new(*campaign.ServiceImpl)),
	campaign.NewHandler,
	wire.Bind(new(campaign.Controller), new(*campaign.Handler)),
)

func InitializeUserController(gormConnection *gorm.DB, validatorInstance *validator.Validate, viperConfig *viper.Viper) user.Controller {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeCampaignController(gormConnection *gorm.DB, validatorInstance *validator.Validate) campaign.Controller {
	wire.Build(campaignFeatureSet)
	return nil
}
