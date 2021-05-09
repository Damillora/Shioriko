package services

import (
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/google/uuid"
)

const perPage = 20

func GetPostAll(page int) []database.Post {
	var posts []database.Post
	database.DB.Joins("Blob").Preload("Tags").Preload("Tags.TagType").Offset((page - 1) * perPage).Limit(20).Find(&posts)
	return posts
}

func GetPostTag(page int, tagSyntax string) []database.Post {
	tag, err := GetTag(tagSyntax)
	if err != nil {
		return []database.Post{}
	}
	var posts []database.Post
	database.DB.Model(&tag).Joins("Blob").Preload("Tags").Preload("Tags.TagType").Offset((page - 1) * perPage).Limit(20).Association("Posts").Find(&posts)
	return posts
}

func GetPost(id string) (*database.Post, error) {
	var post database.Post
	result := database.DB.Joins("Blob").Preload("Tags").Preload("Tags.TagType").Where("posts.id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}
func CreatePost(model models.PostCreateModel) (*database.Post, error) {
	tags, err := ParseTags(model.Tags)
	if err != nil {
		return nil, err
	}

	post := database.Post{
		ID:        uuid.NewString(),
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

	var post database.Post
	result := database.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	post.SourceURL = model.SourceURL
	post.Tags = tags

	result = database.DB.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil

}
func CountPostPages() int {
	var count int64
	database.DB.Model(&database.Post{}).Count(&count)
	return int(count/perPage) + 1
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
