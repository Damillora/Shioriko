package app

import (
	"net/http"
	"strconv"

	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializeTagTypeRoutes(g *gin.Engine) {
	unprotected := g.Group(("/api/tagtype"))
	{
		unprotected.GET("/", tagTypeGet)
	}
	protected := g.Group("/api/tagtype").Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", tagTypeCreate)
		protected.DELETE("/:id", tagTypeDelete)
	}
}

func tagTypeGet(c *gin.Context) {
	tagTypes := services.GetTagTypeAll()

	var tagResult []models.TagTypeListItem

	for _, tagType := range tagTypes {
		tagResult = append(tagResult, models.TagTypeListItem{
			ID:   tagType.ID,
			Name: tagType.Name,
		})
	}
	c.JSON(http.StatusOK, tagResult)
}

func tagTypeCreate(c *gin.Context) {
	var model models.TagTypeCreateModel
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		c.Abort()
		return
	}

	tagType, err := services.CreateOrUpdateTagType(model)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, models.TagTypeListItem{
		ID:   tagType.ID,
		Name: tagType.Name,
	})

}

func tagTypeDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = services.DeleteTagType(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, models.ErrorResponse{
		Code:    http.StatusOK,
		Message: "Success",
	})
}
