package services

import (
	"log"
	"math"

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
	log.Println(tags)
	if err != nil {
		return []database.Post{}
	}
	var posts []database.Post
	database.DB.Model(&tags).Distinct().Joins("Blob").Preload("Tags").Preload("Tags.TagType").Order("created_at desc").Offset((page - 1) * perPage).Limit(20).Association("Posts").Find(&posts)
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
	result := database.DB.Preload("Tags").Where("id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	database.DB.Model(&post).Association("Tags").Replace(tags)

	post.SourceURL = model.SourceURL

	result = database.DB.Save(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil

}
func CountPostPages() int {
	var count int64
	database.DB.Model(&database.Post{}).Count(&count)
	return int(math.Abs(float64(count-1))/perPage) + 1
}
func CountPostPagesTag(tagSyntax []string) int {
	tags, err := ParseReadTags(tagSyntax)
	if err != nil {
		return 0
	}

	var count int64
	count = database.DB.Model(&tags).Distinct().Joins("Blob").Preload("Tags").Preload("Tags.TagType").Association("Posts").Count()

	return int(math.Abs(float64(count-1))/perPage) + 1
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
