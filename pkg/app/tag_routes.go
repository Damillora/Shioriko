package app

import (
	"net/http"

	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
)

func InitializeTagRoutes(g *gin.Engine) {
	autocomplete := g.Group("/api/tag-autocomplete")
	{
		autocomplete.GET("/", tagAutocomplete)
	}
	unprotected := g.Group("/api/tag")
	{
		unprotected.GET("/", tagGet)
		unprotected.GET("/:tag", tagGetOne)
	}
	related := g.Group("/api/tag-related")
	{
		related.GET("/:tag", tagGetRelated)
	}
	protected := g.Group("/api/tag").Use(middleware.AuthMiddleware())
	{
		protected.PUT("/:tag/note", tagUpdateNote)
		protected.PUT("/:tag", tagUpdate)
	}
}

func tagGet(c *gin.Context) {
	tags := services.GetTagAll()
	c.JSON(http.StatusOK, tags)
}

func tagGetOne(c *gin.Context) {
	tag := c.Param("tag")
	tagObj, err := services.GetTag(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, models.TagReadModel{
		TagID:     tagObj.TagID,
		TagName:   tagObj.TagName,
		TagType:   tagObj.TagType,
		TagNote:   tagObj.TagNote,
		PostCount: tagObj.PostCount,
	})
}

func tagGetRelated(c *gin.Context) {
	tag := c.Param("tag")
	tagObj, err := services.GetRelatedTags(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, tagObj)
}

func tagAutocomplete(c *gin.Context) {
	tags := services.GetTagAutocomplete()
	c.JSON(http.StatusOK, tags)
}

func tagUpdateNote(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User don't exist",
		})
		c.Abort()
	}

	var model models.TagNoteUpdateModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	tag := c.Param("tag")
	err = services.UpdateTagNotes(tag, model.Note)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, nil)
}

func tagUpdate(c *gin.Context) {
	_, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "User don't exist",
		})
		c.Abort()
	}

	var model models.TagUpdateModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	tag := c.Param("tag")
	err = services.UpdateTag(tag, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, nil)
}
