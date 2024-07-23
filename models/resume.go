package models

import (
	"github.com/jinzhu/gorm"
)

type Resume struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	FilePath string `json:"file_path"`
}
