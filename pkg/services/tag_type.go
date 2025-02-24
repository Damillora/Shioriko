package services

import (
	"github.com/Damillora/phoebe/pkg/database"
	"github.com/Damillora/phoebe/pkg/models"
)

func GetTagTypeAll() []database.TagType {
	var tagTypes []database.TagType
	database.DB.Find(&tagTypes)
	return tagTypes
}

func CreateOrUpdateTagType(model models.TagTypeCreateModel) (*database.TagType, error) {
	var tagType database.TagType
	result := database.DB.Where("name = ?", model.Name).First(&tagType)

	if result.Error != nil {
		tagType = database.TagType{
			Name: model.Name,
		}
		result = database.DB.Create(&tagType)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return &tagType, nil
}

func DeleteTagType(id uint) error {
	var tagType database.TagType
	database.DB.Where("id = ?", id).First(&tagType)
	result := database.DB.Delete(tagType)
	return result.Error
}
