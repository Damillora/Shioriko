package services

import (
	"errors"
	"strings"

	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/google/uuid"
)

func GetTagAll() []database.Tag {
	var tags []database.Tag
	database.DB.Joins("TagType").Find(&tags)
	return tags
}

func CreateOrUpdateTagGeneric(tagName string) (*database.Tag, error) {
	var tag database.Tag
	result := database.DB.Where("name = ?", tagName).First(&tag)
	if result.Error != nil {
		var tagType database.TagType
		database.DB.Where("name = ?", "general").First(&tagType)

		tag = database.Tag{
			ID:        uuid.NewString(),
			Name:      tagName,
			TagTypeID: tagType.ID,
		}
		result = database.DB.Create(&tag)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return &tag, nil
}
func CreateOrUpdateTagComplex(tagName string, tagTypeString string) (*database.Tag, error) {
	var tag database.Tag
	var tagType database.TagType
	result := database.DB.Where("name = ?", tagTypeString).First(&tagType)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.DB.Where("name = ? AND tag_type_id = ? ", tagName, tagType.ID).First(&tag)

	if result.Error != nil {
		tag = database.Tag{
			ID:        uuid.NewString(),
			Name:      tagName,
			TagTypeID: tagType.ID,
		}
		result = database.DB.Create(&tag)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return &tag, nil
}
func CreateOrUpdateTag(tagSyntax string) (*database.Tag, error) {
	tagFields := strings.Split(tagSyntax, ":")
	var tagName string
	var tagType string
	if len(tagFields) == 1 {
		tagName = tagFields[0]
		return CreateOrUpdateTagGeneric(tagName)
	} else if len(tagFields) == 2 {
		tagType = tagFields[0]
		tagName = tagFields[1]
		return CreateOrUpdateTagComplex(tagName, tagType)
	} else {
		return nil, errors.New("Malformed tag syntax")
	}
}

func GetTag(tagSyntax string) (*database.Tag, error) {
	tagFields := strings.Split(tagSyntax, ":")
	var tagName string
	var tagType database.TagType
	if len(tagFields) == 1 {
		tagName = tagFields[0]
		database.DB.Where("name = ?", "general").First(&tagType)
	} else if len(tagFields) == 2 {
		tagName = tagFields[1]
		result := database.DB.Where("name = ?", tagFields[0]).First(&tagType)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		return nil, errors.New("Malformed tag syntax")
	}
	var tag database.Tag
	result := database.DB.Preload("Posts").Where("name = ? AND tag_type_id = ? ", tagName, tagType.ID).First(&tag)

	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func ParseTags(tags []string) ([]database.Tag, error) {
	var result []database.Tag
	for _, tagSyntax := range tags {
		tag, err := CreateOrUpdateTag(tagSyntax)
		if err != nil {
			return nil, err
		}
		result = append(result, *tag)
	}
	return result, nil
}
