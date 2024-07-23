package main

import (
	"assignment/controller"
	"assignment/database"
	"assignment/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectToDB()

	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/admin/job/:job_id", controller.GetJob)
		auth.GET("/jobs", controller.GetJobs)
		auth.POST("/admin/job", controller.CreateJob)
		auth.POST("/uploadresume", controller.UploadResume)
		auth.GET("/admin/applicant/:applicant_id", controller.GetApplicantData)
		auth.GET("/admin/applicants", controller.GetApplicants)

	}

	r.Run(":8080")
}
