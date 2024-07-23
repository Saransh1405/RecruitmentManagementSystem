package controller

import (
	"assignment/database"
	"assignment/models"
	"assignment/util"
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadResume(c *gin.Context) {
	file, _, err := c.Request.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(file)
	fileBytes := buffer.Bytes()

	resp, err := util.UploadResumeToAPI(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// Assume you get user ID from context
	userID := c.GetInt("user_id")

	resume := models.Resume{
		UserID:   uint(userID),
		FilePath: resp.FilePath,
	}

	if err := database.DB.Create(&resume).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Resume uploaded"})
}
