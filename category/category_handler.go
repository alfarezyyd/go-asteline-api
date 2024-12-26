package category

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/category/dto"
	"go-asteline-api/helper"
	"net/http"
)

type Handler struct {
	categoryService Service
}

func NewHandler(categoryService Service) *Handler {
	return &Handler{
		categoryService: categoryService,
	}
}

func (categoryHandler *Handler) GetAll(ginContext *gin.Context) {
	allCategoryModel := categoryHandler.categoryService.GetAll(ginContext)
	ginContext.JSON(http.StatusOK, allCategoryModel)
}

func (categoryHandler *Handler) Create(ginContext *gin.Context) {
	var categoryCreateDto dto.CategoryCreateDto
	err := ginContext.ShouldBindJSON(&categoryCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	categoryHandler.categoryService.HandleCreate(ginContext, &categoryCreateDto)
}

func (categoryHandler *Handler) Update(ginContext *gin.Context) {
	var categoryUpdateDto dto.CategoryUpdateDto
	err := ginContext.ShouldBindJSON(&categoryUpdateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	categoryHandler.categoryService.HandleUpdate(ginContext, &categoryUpdateDto)
	ginContext.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func (categoryHandler *Handler) Delete(ginContext *gin.Context) {

}
