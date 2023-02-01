package main

import (
	"go-api/controllers"
	"go-api/initializers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	initializers.InitializerEnvVariable()
	initializers.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.GET("/posts/page", controllers.GetPostPerPage)
	r.Run()
}
