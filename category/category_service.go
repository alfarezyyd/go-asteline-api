package category

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/category/dto"
)

type Service interface {
	GetAll(ginContext *gin.Context)
	HandleCreate(ginContext *gin.Context, categoryCreateDto *dto.CategoryCreateDto)
	HandleUpdate(ginContext *gin.Context, categoryUpdateDto *dto.CategoryUpdateDto)
	HandleDelete(ginContext *gin.Context)
}
