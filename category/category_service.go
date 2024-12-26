package category

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/category/dto"
	"go-asteline-api/model"
)

type Service interface {
	GetAll(ginContext *gin.Context) []model.Category
	HandleCreate(ginContext *gin.Context, categoryCreateDto *dto.CategoryCreateDto)
	HandleUpdate(ginContext *gin.Context, categoryUpdateDto *dto.CategoryUpdateDto)
	HandleDelete(ginContext *gin.Context)
}
