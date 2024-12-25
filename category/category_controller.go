package category

import "github.com/gin-gonic/gin"

type Controller interface {
	GetAll(ginContext *gin.Context)
	Create(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
}
