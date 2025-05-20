package controller

import (
	"ginAndgo_zore/model"
	"ginAndgo_zore/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type PostController struct {
	Db *gorm.DB
}

func (pc *PostController) PostsAdd(c *gin.Context) {

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(util.BindError.Code, gin.H{"error": util.BindError.Msg})
	}

	post.CreatedAt = time.Now().Local()
	if err := pc.Db.Create(&post).Error; err != nil {
		c.JSON(util.CreateError.Code, gin.H{"error": util.CreateError.Msg})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (pc *PostController) GetPost(c *gin.Context) {

	var post model.Post
	postid, _ := strconv.Atoi(c.Param("postid"))
	if err := pc.Db.Model(&model.Post{}).Where("post_id=?", postid).First(&post).Error; err != nil {
		c.JSON(util.DBError.Code, gin.H{"error": util.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}
func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []model.Post
	if err := pc.Db.Find(&posts).Error; err != nil {
		c.JSON(util.DBError.Code, gin.H{"error": util.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(util.BindError.Code, gin.H{"error": util.BindError.Msg})
		return
	}
	post.Id, _ = strconv.Atoi(c.Param("postid"))
	post.UpdatedAt = time.Now().Local()
	if err := pc.Db.Model(&model.Post{}).Where("post_id", post.Id).Updates(post).Error; err != nil {
		c.JSON(util.DBError.Code, gin.H{"error": util.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func (pc *PostController) DeletePost(c *gin.Context) {
	var post model.Post
	post.Id, _ = strconv.Atoi(c.Param("postid"))
	if err := pc.Db.Model(&model.Post{}).Where("post_id=?", post.Id).Delete(&post).Error; err != nil {
		c.JSON(util.DBError.Code, gin.H{"error": util.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, post)
	}
}
