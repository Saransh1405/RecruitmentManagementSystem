package controller

import (
	"assignment/database"
	"assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	var job models.Job

	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job created successfully",
	})
}

func GetJob(c *gin.Context) {
	var job models.Job

	jobID := c.Param("job_id")

	if err := database.DB.Preload("Applicants").First(&job, jobID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No job found",
		})
	}

	c.JSON(http.StatusOK, job)
}

func GetJobs(c *gin.Context) {
	var jobs []models.Job
	if err := database.DB.Preload("Applicants").Find(&jobs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No jobs found",
		})
	}
	c.JSON(http.StatusOK, jobs)

}
