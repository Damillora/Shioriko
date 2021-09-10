package services

import (
	"errors"
	"strings"

	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/google/uuid"
)

func GetTagAll() []models.TagListItem {
	var tags []models.TagListItem
	database.DB.Model(&database.Tag{}).
		Joins("join tag_types on tag_types.id = tags.tag_type_id").
		Joins("left join post_tags on post_tags.tag_id = tags.id").
		Select("tags.id as tag_id, tags.name as tag_name, tag_types.name as tag_type, count(post_tags.post_id) as post_count").
		Group("tags.id, tags.name, tag_types.name").
		Order("post_count DESC").
		Find(&tags)
	return tags
}

func GetTagAutocomplete() []string {
	var tags []string
	result := database.DB.Model(&database.Tag{}).
		Joins("join tag_types on tag_types.id = tags.tag_type_id").
		Select("concat(tag_types.name,':',tags.name) as name").
		Find(&tags)
	if result.Error != nil {
		return []string{}
	}
	return tags
}

func FindTagGeneric(tagName string) (*database.Tag, error) {
	var tag database.Tag
	result := database.DB.Where("name = ?", tagName).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}
func FindTagComplex(tagName string, tagTypeString string) (*database.Tag, error) {
	var tag database.Tag
	var tagType database.TagType
	result := database.DB.Where("name = ?", tagTypeString).First(&tagType)
	if result.Error != nil {
		return nil, result.Error
	}

	result = database.DB.Where("name = ? AND tag_type_id = ? ", tagName, tagType.ID).First(&tag)

	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func CreateOrUpdateTagGeneric(tagName string) (*database.Tag, error) {
	tag, err := FindTagGeneric(tagName)
	if err != nil {
		var tagType database.TagType
		database.DB.Where("name = ?", "general").First(&tagType)

		tag = &database.Tag{
			ID:        uuid.NewString(),
			Name:      tagName,
			TagTypeID: tagType.ID,
		}
		result := database.DB.Create(&tag)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return tag, nil
}
func CreateOrUpdateTagComplex(tagName string, tagTypeString string) (*database.Tag, error) {
	tag, err := FindTagComplex(tagName, tagTypeString)

	if err != nil {
		var tagType database.TagType
		result := database.DB.Where("name = ?", tagTypeString).First(&tagType)
		if result.Error != nil {
			return nil, result.Error
		}
		tag = &database.Tag{
			ID:        uuid.NewString(),
			Name:      tagName,
			TagTypeID: tagType.ID,
		}
		result = database.DB.Create(&tag)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return tag, nil
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
func FindTag(tagSyntax string) (*database.Tag, error) {
	tagFields := strings.Split(tagSyntax, ":")
	var tagName string
	var tagType string
	if len(tagFields) == 1 {
		tagName = tagFields[0]
		return FindTagGeneric(tagName)
	} else if len(tagFields) == 2 {
		tagType = tagFields[0]
		tagName = tagFields[1]
		return FindTagComplex(tagName, tagType)
	} else {
		return nil, errors.New("Malformed tag syntax")
	}
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

func ParseReadTags(tags []string) ([]database.Tag, error) {
	var result []database.Tag
	for _, tagSyntax := range tags {
		tag, err := FindTag(tagSyntax)
		if err != nil {
			return nil, err
		}
		result = append(result, *tag)
	}
	return result, nil
}
