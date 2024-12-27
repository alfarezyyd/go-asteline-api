package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/spf13/viper"
	"go-asteline-api/config"
	"go-asteline-api/exception"
	"go-asteline-api/middleware"
	"go-asteline-api/routes"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	viperConfig := viper.New()
	viperConfig.SetConfigFile(".env")
	viperConfig.AddConfigPath(".")
	viperConfig.AutomaticEnv()
	viperConfig.ReadInConfig()

	// Load Google OAUTH2 Credentials
	googleClientId := viperConfig.GetString("GOOGLE_CLIENT_ID")
	googleClientSecret := viperConfig.GetString("GOOGLE_CLIENT_SECRET")
	googleClientCallbackUrl := viperConfig.GetString("GOOGLE_CLIENT_CALLBACK_URL")
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, googleClientCallbackUrl),
	)

	// Initialize
	ginEngine := gin.Default()
	databaseInstance := config.NewDatabaseConnection()
	databaseConnection := databaseInstance.GetDatabaseConnection()
	validatorInstance := config.InitializeValidator()
	midtrans.ServerKey = viperConfig.GetString("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
	var midtransCoreClient = coreapi.Client{}
	midtransCoreClient.New(midtrans.ServerKey, midtrans.Sandbox)
	// Interceptor
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(exception.Interceptor())

	// Injection of Dependency
	userController := InitializeUserController(databaseConnection, validatorInstance, viperConfig)
	campaignController := InitializeCampaignController(databaseConnection, validatorInstance)
	donationController := InitializeDonationController(databaseConnection, validatorInstance, &midtransCoreClient)
	categoryController := InitializeCategoryController(databaseConnection, validatorInstance)
	routes.PublicRoute(ginEngine, userController, campaignController, donationController)
	apiRouterGroup := ginEngine.Group("/api")
	apiRouterGroup.Use(middleware.AuthMiddleware(viperConfig))
	routes.UserRoute(apiRouterGroup, campaignController, categoryController)
	err := ginEngine.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
