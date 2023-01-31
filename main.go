package main

import (
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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "hello"})
	})
	r.Run()
}
