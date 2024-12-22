package main

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/config"
	"go-asteline-api/routes"
	"go-asteline-api/user"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// Initialize
	ginEngine := gin.Default()
	// Database
	databaseInstance := config.NewDatabaseConnection()
	databaseConnection := databaseInstance.GetDatabaseConnection()

	// Routes
	ginEngine.Group("/api")

	// Injection of User
	userRepository := user.NewRepository()
	userService := user.NewService(userRepository, databaseConnection)
	userController := user.NewHandler(&userService)
	routes.UserRoute(ginEngine, userController)

	err := ginEngine.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
