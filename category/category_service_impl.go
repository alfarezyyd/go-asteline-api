package category

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-asteline-api/category/dto"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	categoryRepository Repository
	dbConnection       *gorm.DB
	validatorInstance  *validator.Validate
}

func NewService(categoryRepository Repository, dbConnection *gorm.DB, validatorInstance *validator.Validate) *ServiceImpl {
	return &ServiceImpl{
		categoryRepository: categoryRepository,
		dbConnection:       dbConnection,
		validatorInstance:  validatorInstance,
	}
}

func (categoryService *ServiceImpl) GetAll(c *gin.Context) {}

func (categoryService *ServiceImpl) HandleCreate(ginContext *gin.Context, categoryCreateDto *dto.CategoryCreateDto) {
	err := categoryService.validatorInstance.Struct(categoryCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	categoryModel, err := mapper.MapCategoryDtoIntoCategoryModel(categoryCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	categoryService.dbConnection.Create(categoryModel)
	ginContext.JSON(http.StatusCreated, categoryModel)
}

func (categoryService *ServiceImpl) HandleUpdate(ginContext *gin.Context, categoryUpdateDto *dto.CategoryUpdateDto) {
}

func (categoryService *ServiceImpl) HandleDelete(ginContext *gin.Context) {
}
