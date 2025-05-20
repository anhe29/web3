package controller

import (
	"ginAndgo_zore/config"
	"ginAndgo_zore/model"
	"ginAndgo_zore/util"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Db     *gorm.DB
	Config *config.Config
}

func (DB *AuthController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(util.BindError.Code, gin.H{"error": util.BindError.Msg})
	}
	//加密密码
	bcyPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(util.HashError.Code, gin.H{"error": util.HashError.Msg})
		return
	}
	user.Password = string(bcyPassword)
	if err := DB.Db.Create(&user).Error; err != nil {
		c.JSON(util.CreateError.Code, gin.H{"error": util.CreateError.Msg})
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}

// 登录
func (DB *AuthController) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(util.BindError.Code, gin.H{"error": util.BindError.Msg})
	}
	var userInfo model.User
	if err := DB.Db.Where("username = ?", user.Username).First(&userInfo).Error; err != nil {
		c.JSON(util.UsernameOrPasswordError.Code, gin.H{"error": util.UsernameOrPasswordError.Msg})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password)); err != nil {
		c.JSON(util.UsernameOrPasswordError.Code, gin.H{"error": util.UsernameOrPasswordError.Msg})
	}

	DB.Config = config.LoadConfig()
	token, err := util.GenerateToken(userInfo, DB.Config.JWTSecret)
	if err != nil {
		c.JSON(util.GenerateTokenError.Code, gin.H{"error": util.GenerateTokenError.Msg})
		return
	}
	c.Header("Authorization", "Bearer"+token)
}
