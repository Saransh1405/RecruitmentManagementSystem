package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Email           string `json:"email" gorm:"unique"`
	Address         string `json:"address"`
	UserType        string `json:"user_type"` // "Admin" or "Applicant"
	PasswordHash    string `json:"password_hash"`
	ProfileHeadline string `json:"profile_headline"`
	Profile         string `json:"profile"`
}

type Profile struct {
	gorm.Model
	UserID            uint   `json:"user_id"`
	ResumeFileAddress string `json:"resume_file_address"`
	Skills            string `json:"skills"`
	Education         string `json:"education"`
	Experience        string `json:"experience"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
}
