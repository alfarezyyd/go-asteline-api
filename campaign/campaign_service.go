package campaign

import "github.com/gin-gonic/gin"

type Service interface {
	HandleCreate(ginContext *gin.Context)
}
