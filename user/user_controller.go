package user

import "github.com/gin-gonic/gin"

type Controller interface {
	Login(ginContext *gin.Context)
	Register(ginContext *gin.Context)
	LoginWithProvider(ginContext *gin.Context)
	ProviderCallback(ginContext *gin.Context)
	LoginProviderSuccess(ginContext *gin.Context)
}
