package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-asteline-api/category/dto"
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
	categoryId := ginContext.Param("id")
	if categoryId == "" {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}
	err := categoryService.validatorInstance.Struct(categoryUpdateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	categoryModel, err := mapper.MapCategoryDtoIntoCategoryModel(categoryUpdateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	var existingCategoryModel model.Category
	gormTransaction := categoryService.dbConnection.Begin()
	gormTransaction.Where("id = ?", categoryId).First(&existingCategoryModel)
	err = mapper.MapExistingModelIntoUpdateModel(*categoryUpdateDto, *categoryModel)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
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
	dbConn := gormTransaction.Where("id = ?", categoryId).Delete(&model.Category{})
	if helper.CheckErrorOperation(dbConn.Error, ginContext, http.StatusBadRequest) {
		return
	}
	helper.TransactionOperation(gormTransaction, ginContext)
	ginContext.JSON(http.StatusOK, gin.H{"message": "ok"})
}
