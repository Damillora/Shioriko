package app

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitializePostRoutes(g *gin.Engine) {
	unprotected := g.Group("/api/post")
	{
		unprotected.GET("/", postGet)
		unprotected.GET("/:id", postGetOne)
	}
	protected := g.Group("/api/post").Use(middleware.AuthMiddleware())
	{
		protected.POST("/create", postCreate)
		protected.POST("/:id", postUpdate)
		protected.DELETE("/:id", postDelete)
	}

}

func postGet(c *gin.Context) {
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	tag := c.Query("tags")

	tags := strings.Split(tag, " ")

	var posts []database.Post
	var postPages int

	if tag != "" {
		posts = services.GetPostTags(page, tags)
		postPages = services.CountPostPagesTag(tags)
	} else {
		posts = services.GetPostAll(page)
		postPages = services.CountPostPages()
	}

	var postResult []models.PostListItem
	for _, post := range posts {
		var tagStrings []string
		for _, tag := range post.Tags {
			tagStrings = append(tagStrings, tag.TagType.Name+":"+tag.Name)
		}

		postResult = append(postResult, models.PostListItem{
			ID:        post.ID,
			ImagePath: "/data/" + post.Blob.FilePath,
			Tags:      tagStrings,
		})
	}
	c.JSON(http.StatusOK, models.PostPaginationResponse{
		CurrentPage: page,
		PostCount:   postPages,
		Posts:       postResult,
	})
}

func postGetOne(c *gin.Context) {
	id := c.Param("id")
	post, err := services.GetPost(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	var tagStrings []string
	for _, tag := range post.Tags {
		tagStrings = append(tagStrings, tag.TagType.Name+":"+tag.Name)
	}

	c.JSON(http.StatusOK, models.PostReadModel{
		ID:        post.ID,
		ImagePath: "/data/" + post.Blob.FilePath,
		SourceURL: post.SourceURL,
		Tags:      tagStrings,
	})
}

func postCreate(c *gin.Context) {
	var model models.PostCreateModel
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

	post, err := services.CreatePost(model)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, models.PostReadModel{
		ID:        post.ID,
		ImagePath: "/data/" + post.Blob.FilePath,
	})
}

func postUpdate(c *gin.Context) {
	id := c.Param("id")

	var model models.PostUpdateModel
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

	post, err := services.UpdatePost(id, model)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, models.PostListItem{
		ID:        post.ID,
		ImagePath: "/data/" + post.Blob.FilePath,
	})

}

func postDelete(c *gin.Context) {
	id := c.Param("id")

	err := services.DeletePost(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ErrorResponse{
		Code:    http.StatusOK,
		Message: "Success",
	})

}
