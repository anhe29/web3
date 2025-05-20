package main

import (
	"ginAndgo_zore/config"
	"ginAndgo_zore/controller"
	"ginAndgo_zore/middlewares"
	"ginAndgo_zore/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db := util.Conn()
	DB = db
}

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()
	userController := controller.AuthController{Db: DB, Config: cfg}
	postController := controller.PostController{Db: DB}
	commentController := controller.CommentController{Db: DB}
	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
		public.GET("/posts/:postid", postController.GetPost)
		public.GET("/posts", postController.GetPosts)
		public.GET("/posts/:postid/comments", commentController.GetComments)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuth(cfg))
	{
		protected.POST("/posts/add", postController.PostsAdd)
		protected.PUT("/posts/:postid", postController.UpdatePost)
		protected.DELETE("/posts/:postid", postController.DeletePost)
		protected.POST("/posts/:postid/comments", commentController.CreateComment)
	}

	err := r.Run(":8080") //要放在最后面
	if err != nil {
		panic(err)
	}
}
