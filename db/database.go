package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("go_api.db"))

	if err != nil {
		fmt.Printf("error init database %v", err.Error())
		return
	}

	Migrations(db)

	Database = db
}
