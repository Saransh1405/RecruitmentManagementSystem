package database

import (
	"assignment/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := "postgresql://assignment_2yy2_user:tcyRLggCMpjAwaOYGEhYdceOBstbEdeJ@dpg-cqfmu5iju9rs73c0414g-a.oregon-postgres.render.com/assignment_2yy2"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	DB.AutoMigrate(&models.User{}, &models.Job{}, &models.Resume{})
}
