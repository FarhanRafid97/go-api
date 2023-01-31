package controllers

import (
	"fmt"
	"go-api/initializers"
	"go-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	query := c.Request.URL.Query()
	fmt.Println(query)
	post := models.Post{Title: "hello", Body: "world"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": post})
}

func GetPost(c *gin.Context) {
	post := []models.Post{}
	query := c.Request.URL.Query()
	limit := query.Get("limit")
	cursor := query.Get("cursor")
	cursorNum, _ := strconv.Atoi(cursor)
	limitNUm, _ := strconv.Atoi(limit)
	limitPlusONe := limitNUm + 1

	isMore := true

	result := initializers.DB.Where("ID > ? ", cursorNum).Limit(limitPlusONe).Find(&post)
	if len(post) == limitPlusONe {
		isMore = true
	} else {
		isMore = false
	}
	if result.Error != nil {
		c.Status(400)
		return
	}
	fmt.Println(len(post))
	post = post[:len(post)-1]
	fmt.Println(len(post))

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": post, "IsMore": isMore})
}
