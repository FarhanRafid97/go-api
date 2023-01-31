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
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPost)
	r.Run()
}
