package app

import (
	"net/http"
	"path/filepath"

	"github.com/Damillora/Shioriko/pkg/config"
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitializeBlobRoutes(g *gin.Engine) {
	protected := g.Group("/api/blob").Use(middleware.AuthMiddleware())
	{
		protected.POST("/upload", uploadBlob)
	}

}

func uploadBlob(c *gin.Context) {
	dataDir := config.CurrentConfig.DataDirectory
	// Source
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	id := uuid.NewString()
	filename := id + filepath.Ext(file.Filename)

	err = c.SaveUploadedFile(file, filepath.Join(dataDir, filename))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	blob := database.Blob{
		ID:       id,
		FilePath: filename,
	}
	database.DB.Create(&blob)

	c.JSON(http.StatusOK, models.BlobResponse{
		ID: id,
	})
}
