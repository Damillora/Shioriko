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

func InitializePostRoutes(g *gin.Engine) {
	unprotected := g.Group("/api/post")
	{
		unprotected.GET("/", postGet)
		unprotected.GET("/:id", postGetOne)
		unprotected.GET("/tag/:id", postGetTag)
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
	posts := services.GetPostAll(page)
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
	postPages := services.CountPostPages()
	c.JSON(http.StatusOK, models.PostPaginationResponse{
		CurrentPage: page,
		TotalPage:   postPages,
		Posts:       postResult,
	})
}

func postGetTag(c *gin.Context) {
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	tag := c.Param("id")

	posts := services.GetPostTag(page, tag)
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
	postPages := services.CountPostPages()
	c.JSON(http.StatusOK, models.PostPaginationResponse{
		CurrentPage: page,
		TotalPage:   postPages,
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
