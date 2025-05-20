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

type CommentController struct {
	Db *gorm.DB
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(util.BindError.Code, gin.H{"error": util.BindError.Msg})
		return
	}
	comment.UserId = c.MustGet("user_id").(int)
	comment.PostId, _ = strconv.Atoi(c.Param("postid"))
	comment.CreatedAt = time.Now().Local()
	if err := cc.Db.Create(&comment).Error; err != nil {
		c.JSON(util.CreateError.Code, gin.H{"error": util.CreateError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

func (cc *CommentController) GetComments(c *gin.Context) {
	var comments []model.Comment
	postid, _ := strconv.Atoi(c.Param("postid"))
	if err := cc.Db.Model(&model.Comment{}).Where("post_id=?", postid).Find(&comments).Error; err != nil {
		c.JSON(util.DBError.Code, gin.H{"error": util.DBError.Msg})
		return
	} else {
		c.JSON(http.StatusOK, comments)
	}
}
