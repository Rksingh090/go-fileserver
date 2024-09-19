package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"fileserver/config"
	"fileserver/i18n"
	"fileserver/models"
	"fileserver/util"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// Retrieve form field
	file, err := c.FormFile("file")
	upload_directory := c.PostForm("directory")
	used_for := c.PostForm("usage")

	// Retrieve the file from the form
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  i18n.StatusError,
		})
		c.Abort()
		return
	}

	// Generate a unique file name
	uniqueFileName := util.GenerateUniqueFileName(file.Filename)

	// Determine the file type and set the subdirectory
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	subDir := ""
	switch fileExt {
	case ".jpg", ".jpeg", ".png":
		subDir = "images"
	case ".mp4", ".avi":
		subDir = "videos"
	default:
		subDir = "others"
	}

	// Construct the full path to save the file and db
	baseDir := filepath.Join(subDir, upload_directory, uniqueFileName)
	dstDir := filepath.Join(config.FILES_STATIC_DIR, subDir, upload_directory)
	dst := filepath.Join(dstDir, uniqueFileName)
	url := filepath.Join("files", "aplemanvadhikar", baseDir)

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": i18n.PermErrorCreatingDirectory,
			"status":  i18n.StatusError,
		})
		c.Abort()
		return
	}

	imageData := &models.ImageModel{
		FileName:         uniqueFileName,
		OriginalFileName: file.Filename,
		UsedFor:          used_for,
		FilePath:         baseDir,
		Url:              url,
	}
	data, err := models.Insert(config.FS_TABLE_NAME, imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": i18n.InsertInDbError,
			"status":  i18n.StatusError,
		})
		c.Abort()
		return
	}

	// Save the file
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": i18n.ErrorUploadFile,
			"status":  i18n.StatusError,
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.FileUploadedSuccess,
		"status":  i18n.StatusSuccess,
		"data": gin.H{
			"file_path": dst,
			"id":        data,
			"url":       url,
		},
	})
}
