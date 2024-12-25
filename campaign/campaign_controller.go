package campaign

import "github.com/gin-gonic/gin"

type Controller interface {
	GetAll(ginCtx *gin.Context)
	Create(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
}
