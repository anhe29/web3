package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	Id    uint `gorm:"primaryKey" json:"id"`
	Name  string
	Email string
	Posts []Post `gorm:"foreignKey:UserId"`
}
type Post struct {
	Id            uint `gorm:"primaryKey" json:"id"`
	UserId        uint `gorm:"not null" json:"user_id"`
	Title         string
	Body          string
	State         string
	CommentCounts int
	Comments      []Comment `gorm:"foreignKey:PostId"`
}
type Comment struct {
	Id     uint `gorm:"primaryKey" json:"id"`
	PostId uint `gorm:"not null" json:"post_id"`
	UserID uint `gorm:"index"`
	Text   string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 自动迁移：确保表存在
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("迁移数据库失败：" + err.Error())
	}

	user, err := GetUserPostsWithComments(db, 1)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("用户不存在")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(user, err)

	var post Post
	FindMaxCommentsPost(db, &post)
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func GetUserPostsWithComments(db *gorm.DB, userId uint) (User, error) {
	var user User
	err := db.Preload("Posts.Comments").
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}
	return user, nil
}

func FindMaxCommentsPost(db *gorm.DB, post *Post) error {
	err := db.Model(&Post{}).Order("comment_counts desc").First(post)
	return err.Error
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (p *Post) AfterCreate(db *gorm.DB) error {
	result := db.Model(&User{}).Where("id=?", p.UserId).Update("post_counts", gorm.Expr("post_counts+?", 1))
	return result.Error
}
func (c *Comment) BeforeDelete(db *gorm.DB) error {
	var post Post
	if err := db.Model(&Post{}).Where("id=?", c.PostId).Find(&post).Error; err != nil {
		return err
	}
	if post.Id == 0 {
		result := db.Model(&Post{}).Where("id=?", c.PostId).Update("state", "无评论")
		return result.Error
	}
	return nil
}
