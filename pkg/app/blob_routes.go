package app

import (
	"encoding/binary"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"

	_ "golang.org/x/image/webp"

	"github.com/Damillora/Shioriko/pkg/config"
	"github.com/Damillora/Shioriko/pkg/database"
	"github.com/Damillora/Shioriko/pkg/middleware"
	"github.com/Damillora/Shioriko/pkg/models"
	"github.com/Damillora/Shioriko/pkg/services"
	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
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

	previewFilename := id + ".jpg"
	previewFilePath := filepath.Join(dataDir, "preview", folder1, folder2, previewFilename)
	thumbnailFilePath := filepath.Join(dataDir, "thumbnail", folder1, folder2, previewFilename)

	fileObj, _ := file.Open()
	originalImage, _, err := image.Decode(fileObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	width := originalImage.Bounds().Dx()
	height := originalImage.Bounds().Dy()

	hash, err := goimagehash.PerceptionHash(originalImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	hashInt := hash.GetHash()

	similarPosts, err := services.SimilaritySearch(hashInt)

	hashSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(hashSlice, hashInt)

	previewImage := imaging.Resize(originalImage, 1000, 0, imaging.Lanczos)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	thumbnailImage := imaging.Resize(originalImage, 300, 0, imaging.Lanczos)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	previewFile, err := os.Create(previewFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	thumbnailFile, err := os.Create(thumbnailFilePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = jpeg.Encode(previewFile, previewImage, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err = jpeg.Encode(thumbnailFile, thumbnailImage, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	filename := id + filepath.Ext(file.Filename)
	filePath := filepath.Join(dataDir, folder1, folder2, filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	blob := database.Blob{
		ID:                id,
		FilePath:          filepath.Join(folder1, folder2, filename),
		PreviewFilePath:   filepath.Join("preview", folder1, folder2, previewFilename),
		ThumbnailFilePath: filepath.Join("thumbnail", folder1, folder2, previewFilename),
		Width:             width,
		Height:            height,
		Hash1:             hashSlice[0:2],
		Hash2:             hashSlice[2:4],
		Hash3:             hashSlice[4:6],
		Hash4:             hashSlice[6:8],
	}

	database.DB.Create(&blob)

	if len(similarPosts) > 0 {
		c.JSON(http.StatusOK,
			models.BlobSimilarResponse{
				ID:      id,
				Width:   width,
				Height:  height,
				Similar: similarPosts,
			})
		return
	} else {
		c.JSON(http.StatusOK, models.BlobResponse{
			ID:     id,
			Width:  width,
			Height: height,
		})
		return
	}
}
