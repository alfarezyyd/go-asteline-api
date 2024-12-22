package user

import "github.com/gin-gonic/gin"

type Interface interface {
	Register(ginContext *gin.Context)
}
