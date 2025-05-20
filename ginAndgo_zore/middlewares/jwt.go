package middlewares

import (
	"ginAndgo_zore/config"
	"ginAndgo_zore/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer")
		if tokenString == "" {
			c.AbortWithStatusJSON(util.TokenError.Code, gin.H{"error": util.TokenError.Msg})
			return
		}
		if claims, err := util.ParseToken(tokenString, config.JWTSecret); err != nil {
			c.AbortWithStatusJSON(util.TokenError.Code, gin.H{"error": util.TokenError.Msg})
			return
		} else {
			c.Set("user_id", claims.Id)
			c.Set("username", claims.Username)
			c.Set("email", claims.Email)
			c.Next()
		}

	}
}
