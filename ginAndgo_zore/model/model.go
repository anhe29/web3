package model

import "time"

type User struct {
	Id       int `gorm:"primary_key;"`
	Username string
	Password string
	Email    string
	Posts    []Post
	Comments []Comment
}

type Post struct {
	Id        int `gorm:"primary_key;"`
	title     string
	Content   string
	UserId    int
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	Id        int `gorm:"primary_key;"`
	Content   string
	UserId    int
	PostId    int
	CreatedAt time.Time
}
