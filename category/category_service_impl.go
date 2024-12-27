package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-asteline-api/category/dto"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/model"
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

func (categoryService *ServiceImpl) GetAll(ginContext *gin.Context) []model.Category {
	var allCategoryModel []model.Category
	categoryService.dbConnection.Find(&allCategoryModel)
	return allCategoryModel
}

func (categoryService *ServiceImpl) HandleCreate(ginContext *gin.Context, categoryCreateDto *dto.CategoryCreateDto) {
	err := categoryService.validatorInstance.Struct(categoryCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	categoryModel, err := mapper.MapCategoryDtoIntoCategoryModel(categoryCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	categoryService.dbConnection.Create(categoryModel)
	ginContext.JSON(http.StatusCreated, categoryModel)
}

func (categoryService *ServiceImpl) HandleUpdate(ginContext *gin.Context, categoryUpdateDto *dto.CategoryUpdateDto) {
	categoryId := ginContext.Param("id")
	if categoryId == "" {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}
	err := categoryService.validatorInstance.Struct(categoryUpdateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	categoryModel, err := mapper.MapCategoryDtoIntoCategoryModel(categoryUpdateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	var existingCategoryModel model.Category
	gormTransaction := categoryService.dbConnection.Begin()
	gormTransaction.Where("id = ?", categoryId).First(&existingCategoryModel)
	err = mapper.MapExistingModelIntoUpdateModel(*categoryUpdateDto, *categoryModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))

	fmt.Println(existingCategoryModel, categoryModel)
	gormTransaction.Where("id = ?", categoryId).Updates(categoryModel)
	helper.TransactionOperation(gormTransaction, ginContext)
}

func (categoryService *ServiceImpl) HandleDelete(ginContext *gin.Context) {
	categoryId := ginContext.Param("id")
	if categoryId == "" {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}
	gormTransaction := categoryService.dbConnection.Begin()
	err := gormTransaction.Where("id = ?", categoryId).Delete(&model.Category{}).Error
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	helper.TransactionOperation(gormTransaction, ginContext)
}
