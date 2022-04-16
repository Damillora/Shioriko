package services

import (
	"fmt"

	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/google/uuid"
)

const perPage = 20

func GetPostAll(page int) []database.Post {
	var posts []database.Post
	database.DB.Joins("Blob").Preload("Tags").Preload("Tags.TagType").Order("created_at desc").Offset((page - 1) * perPage).Limit(20).Find(&posts)
	return posts
}

func GetPostTags(page int, tagSyntax []string) []database.Post {
	tags, err := ParseReadTags(tagSyntax)

	if err != nil {
		return []database.Post{}
	}

	var tagIds []string
	for _, tag := range tags {
		tagIds = append(tagIds, tag.ID)
	}
	fmt.Printf("%v", tagIds)

	var postIds []string
	database.DB.
		Model(&tags).
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", tagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(tagIds)).
		Distinct().
		Find(&postIds)

	var posts []database.Post
	database.DB.
		Joins("Blob").
		Preload("Tags").
		Preload("Tags.TagType").
		Where("posts.id IN ?", postIds).
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(20).
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
	tags, err := ParseReadTags(tagSyntax)
	if err != nil {
		return 0
	}

	var tagIds []string
	for _, tag := range tags {
		tagIds = append(tagIds, tag.ID)
	}
	fmt.Printf("%v", tagIds)

	var count int64
	database.DB.
		Model(&tags).
		Distinct().
		Joins("join post_tags on post_tags.tag_id = tags.id").
		Select("post_tags.post_id").
		Where("post_tags.tag_id IN ?", tagIds).
		Group("post_tags.post_id").
		Having("count(*) = ?", len(tagIds)).
		Count(&count)

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
