package controller

import (
	"assignment/database"
	"assignment/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetApplicants(c *gin.Context) {
	var users []models.User

	// Only fetch applicants
	if err := database.DB.Where("user_type = ?", "Applicant").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetApplicantData(c *gin.Context) {
	applicantID := c.Param("applicant_id")
	var user models.User
	var profile models.Profile

	id, err := strconv.Atoi(applicantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid applicant ID"})
		return
	}

	// Fetch user data
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Applicant not found"})
		return
	}

	// Fetch profile data
	if err := database.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"profile": profile,
	})
}
