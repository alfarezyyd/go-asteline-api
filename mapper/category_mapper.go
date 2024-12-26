package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-asteline-api/category/dto"
	"go-asteline-api/model"
)

func MapCategoryDtoIntoCategoryModel[T *dto.CategoryCreateDto | *dto.CategoryUpdateDto](categorySaveDto T) (*model.Category, error) {
	var modelCategory model.Category
	err := mapstructure.Decode(categorySaveDto, &modelCategory)
	if err != nil {
		return nil, err
	}
	return &modelCategory, nil
}

func MapExistingModelIntoUpdateModel(categoryUpdateDto dto.CategoryUpdateDto, categoryModel model.Category) error {
	return mapstructure.Decode(categoryModel, &categoryUpdateDto)
}
