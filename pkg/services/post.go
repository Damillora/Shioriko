package services

import (
	"fmt"
	"strings"

	"github.com/Damillora/phoebe/pkg/database"
	"github.com/Damillora/phoebe/pkg/models"
	"github.com/google/uuid"
)

func GetPostAll(page int, perPage int) []database.Post {
	var posts []database.Post
	database.DB.Joins("Blob").Preload("Tags").Preload("Tags.TagType").Order("created_at desc").Offset((page - 1) * perPage).Limit(perPage).Find(&posts)
	return posts
}

func GetPostTags(page int, perPage int, tagSyntax []string) []database.Post {
	positiveTagSyntax := []string{}
	negativeTagSyntax := []string{}

	for _, tag := range tagSyntax {
		if strings.HasPrefix(tag, "-") {
			negativeTagSyntax = append(negativeTagSyntax, strings.TrimPrefix(tag, "-"))
		} else {
			positiveTagSyntax = append(positiveTagSyntax, tag)

		}
	}
	positiveTags, err := ParseReadTags(positiveTagSyntax)
	if err != nil {
		return []database.Post{}
	}
	negativeTags, err := ParseReadTags(negativeTagSyntax)
	if err != nil {
		return []database.Post{}
	}

	var positiveTagIds []string
	for _, tag := range positiveTags {
		positiveTagIds = append(positiveTagIds, tag.ID)
	}

	var negativeTagIds []string
	for _, tag := range negativeTags {
		negativeTagIds = append(negativeTagIds, tag.ID)
	}

	var positivePostIds []string
	database.DB.
		Model(&positiveTags).
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", positiveTagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(positiveTagIds)).
		Distinct().
		Find(&positivePostIds)
	var negativePostIds []string
	database.DB.
		Model(&positiveTags).
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", negativeTagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(negativeTagIds)).
		Distinct().
		Find(&negativePostIds)

	var posts []database.Post
	query := database.DB.
		Joins("Blob").
		Preload("Tags").
		Preload("Tags.TagType")

	if len(positivePostIds) > 0 && len(negativePostIds) > 0 {
		query = query.
			Where("posts.id IN ? AND posts.id NOT IN ?", positivePostIds, negativePostIds)
	} else if len(positivePostIds) > 0 {
		query = query.
			Where("posts.id IN ?", positivePostIds)
	} else if len(negativePostIds) > 0 {
		query = query.
			Where("posts.id NOT IN ?", negativePostIds)
	}
	query.Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&posts)
	return posts
}

func GetPost(id string) (*database.Post, error) {
	var post database.Post
	result := database.DB.Joins("User").Joins("Blob").Preload("Tags").Preload("Tags.TagType").Where("posts.id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}
func CreatePost(userID string, model models.PostCreateModel) (*database.Post, error) {
	tags, err := ParseTags(model.Tags)
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		database.DB.Where("name = ?", "tagme").Find(&tags)
	}
	post := database.Post{
		ID:        uuid.NewString(),
		UserID:    userID,
		BlobID:    model.BlobID,
		SourceURL: model.SourceURL,
		Tags:      tags,
	}

	result := database.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
func UpdatePost(id string, model models.PostUpdateModel) (*database.Post, error) {
	tags, err := ParseTags(model.Tags)
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		database.DB.Where("name = ?", "tagme").Find(&tags)
	}

	var post database.Post
	result := database.DB.Preload("Tags").Where("id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	post.SourceURL = model.SourceURL

	result = database.DB.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	err = database.DB.Model(&post).Association("Tags").Replace(tags)
	if err != nil {
		return nil, err
	}
	return &post, nil

}
func CountPostPages() int {
	var count int64
	database.DB.Model(&database.Post{}).Count(&count)
	return int(count)
}
func CountPostPagesTag(tagSyntax []string) int {
	positiveTagSyntax := []string{}
	negativeTagSyntax := []string{}

	for _, tag := range tagSyntax {
		if strings.HasPrefix(tag, "-") {
			negativeTagSyntax = append(negativeTagSyntax, strings.TrimPrefix(tag, "-"))
		} else {
			positiveTagSyntax = append(positiveTagSyntax, tag)

		}
	}
	positiveTags, err := ParseReadTags(positiveTagSyntax)
	if err != nil {
		return 0
	}
	negativeTags, err := ParseReadTags(negativeTagSyntax)
	if err != nil {
		return 0
	}

	var positiveTagIds []string
	for _, tag := range positiveTags {
		positiveTagIds = append(positiveTagIds, tag.ID)
	}

	var negativeTagIds []string
	for _, tag := range negativeTags {
		negativeTagIds = append(negativeTagIds, tag.ID)
	}

	var positivePostIds []string
	database.DB.
		Model(&positiveTags).
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", positiveTagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(positiveTagIds)).
		Distinct().
		Find(&positivePostIds)
	var negativePostIds []string
	database.DB.
		Model(&positiveTags).
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", negativeTagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(negativeTagIds)).
		Distinct().
		Find(&negativePostIds)

	var count int64
	query := database.DB.
		Model(&database.Post{})
	if len(positivePostIds) > 0 && len(negativePostIds) > 0 {
		query = query.
			Where("posts.id IN ? AND posts.id NOT IN ?", positivePostIds, negativePostIds)
	} else if len(positivePostIds) > 0 {
		query = query.
			Where("posts.id IN ?", positivePostIds)
	} else if len(negativePostIds) > 0 {
		query = query.
			Where("posts.id NOT IN ?", negativePostIds)
	}
	query.
		Count(&count)
	fmt.Println(count)
	return int(count)
}

func DeletePost(id string) error {

	var post database.Post
	result := database.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Where("id = ?", post.BlobID).Delete(&database.Blob{})
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Delete(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
