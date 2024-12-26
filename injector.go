//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-asteline-api/campaign"
	"go-asteline-api/category"
	"go-asteline-api/donation"
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
var categoryFeatureSet = wire.NewSet(
	category.NewRepository,
	wire.Bind(new(category.Repository), new(*category.RepositoryImpl)),
	category.NewService,
	wire.Bind(new(category.Service), new(*category.ServiceImpl)),
	category.NewHandler,
	wire.Bind(new(category.Controller), new(*category.Handler)),
)

var donationFeatureSet = wire.NewSet(
	donation.NewRepository,
	wire.Bind(new(donation.Repository), new(*donation.RepositoryImpl)),
	donation.NewService,
	wire.Bind(new(donation.Service), new(*donation.ServiceImpl)),
	donation.NewHandler,
	wire.Bind(new(donation.Controller), new(*donation.Handler)),
)

func InitializeUserController(gormConnection *gorm.DB, validatorInstance *validator.Validate, viperConfig *viper.Viper) user.Controller {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeCampaignController(gormConnection *gorm.DB, validatorInstance *validator.Validate) campaign.Controller {
	wire.Build(campaignFeatureSet)
	return nil
}

func InitializeCategoryController(gormConnection *gorm.DB, validatorInstance *validator.Validate) category.Controller {
	wire.Build(categoryFeatureSet)
	return nil
}

func InitializeDonationController(gormConnection *gorm.DB, validatorInstance *validator.Validate) donation.Controller {
	wire.Build(donationFeatureSet)
	return nil
}
