package app

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Damillora/Shioriko/pkg/config"
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
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
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	id := uuid.NewString()
	folder1 := id[0:2]
	folder2 := id[2:4]
	if _, err := os.Stat(filepath.Join(dataDir, folder1)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, folder1), 0755)
	}
	if _, err := os.Stat(filepath.Join(dataDir, folder1, folder2)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, folder1, folder2), 0755)
	}

	if _, err := os.Stat(filepath.Join(dataDir, "preview", dataDir, folder1)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, "preview", folder1), 0755)
	}
	if _, err := os.Stat(filepath.Join(dataDir, "preview", dataDir, folder1, folder2)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, "preview", folder1, folder2), 0755)
	}

	if _, err := os.Stat(filepath.Join(dataDir, "thumbnail", dataDir, folder1)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, "thumbnail", folder1), 0755)
	}
	if _, err := os.Stat(filepath.Join(dataDir, "thumbnail", dataDir, folder1, folder2)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, "thumbnail", folder1, folder2), 0755)
	}

	filename := id + filepath.Ext(file.Filename)
	filePath := filepath.Join(dataDir, folder1, folder2, filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	previewFilename := id + ".webp"
	previewFilePath := filepath.Join(dataDir, "preview", folder1, folder2, previewFilename)
	thumbnailFilePath := filepath.Join(dataDir, "thumbnail", folder1, folder2, previewFilename)

	buffer, err := bimg.Read(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	image := bimg.NewImage(buffer)
	metadata, err := image.Metadata()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	width := metadata.Size.Width
	height := metadata.Size.Height

	previewImage, err := image.Resize(1000, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	thumbnailImage, err := image.Resize(300, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = bimg.Write(previewFilePath, previewImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err = bimg.Write(thumbnailFilePath, thumbnailImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	blob := database.Blob{
		ID:                id,
		FilePath:          filepath.Join(folder1, folder2, filename),
		PreviewFilePath:   filepath.Join("preview", folder1, folder2, previewFilename),
		ThumbnailFilePath: filepath.Join("thumbnail", folder1, folder2, previewFilename),
		Width:             width,
		Height:            height,
	}

	database.DB.Create(&blob)

	c.JSON(http.StatusOK, models.BlobResponse{
		ID:     id,
		Width:  width,
		Height: height,
	})
}
