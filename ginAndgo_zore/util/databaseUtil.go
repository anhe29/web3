package util

import (
	"ginAndgo_zore/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Conn() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	return db
}
