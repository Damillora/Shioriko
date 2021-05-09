package app

import (
	"net/http"

	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
)

func InitializeTagRoutes(g *gin.Engine) {
	unprotected := g.Group("/api/tag")
	{
		unprotected.GET("/", tagGet)
	}
}

func tagGet(c *gin.Context) {
	tags := services.GetTagAll()
	var tagResult []models.TagListItem
	for _, tag := range tags {
		tagResult = append(tagResult, models.TagListItem{
			ID:      tag.ID,
			Name:    tag.Name,
			TagType: tag.TagType.Name,
		})
	}
	c.JSON(http.StatusOK, tagResult)
}
