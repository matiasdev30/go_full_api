package db

import (
	"github.com/matiasdev30/go_api/models"
	"gorm.io/gorm"
)

func Migrations(db * gorm.DB){
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Task{})
}