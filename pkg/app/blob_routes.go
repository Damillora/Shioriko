package app

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
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
		return
	}
	fileObj, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	image, _, err := image.Decode(fileObj)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	width := image.Bounds().Dx()
	height := image.Bounds().Dy()

	id := uuid.NewString()
	folder1 := id[0:2]
	if _, err := os.Stat(filepath.Join(dataDir, folder1)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, folder1), 0755)
	}
	folder2 := id[2:4]
	if _, err := os.Stat(filepath.Join(dataDir, folder1, folder2)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dataDir, folder1, folder2), 0755)
	}

	filename := id + filepath.Ext(file.Filename)

	err = c.SaveUploadedFile(file, filepath.Join(dataDir, folder1, folder2, filename))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	blob := database.Blob{
		ID:       id,
		FilePath: filepath.Join(folder1, folder2, filename),
		Width:    width,
		Height:   height,
	}
	database.DB.Create(&blob)

	c.JSON(http.StatusOK, models.BlobResponse{
		ID:     id,
		Width:  width,
		Height: height,
	})
}
