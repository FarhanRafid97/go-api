package controllers

import (
	"fmt"
	"go-api/initializers"
	"go-api/models"
	"math"
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
	posts := []models.Post{}
	post := models.Post{}
	query := c.Request.URL.Query()
	limit := query.Get("limit")
	cursor := query.Get("cursor")
	cursorNum, _ := strconv.Atoi(cursor)
	limitNUm, _ := strconv.Atoi(limit)
	limitPlusONe := limitNUm + 1

	isMore := true

	var count int64
	initializers.DB.Model(&post).Count(&count)
	fmt.Println(count)
	result := initializers.DB.Where("ID > ? ", cursorNum).Limit(limitPlusONe).Find(&posts)
	if len(posts) == limitPlusONe {
		isMore = true
	} else {
		isMore = false
	}
	if result.Error != nil {
		c.Status(400)
		return
	}

	posts = posts[:len(posts)-1]

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": posts, "IsMore": isMore})
}

func GetPostPerPage(c *gin.Context) {
	posts := []models.Post{}
	post := models.Post{}
	query := c.Request.URL.Query()
	page := query.Get("page")
	pageNum, _ := strconv.Atoi(page)
	currentPage := 1

	if pageNum != 0 {
		currentPage = pageNum
	}

	currentId := 0

	if currentPage == 1 {
		currentId = 0
	} else {
		state := currentPage - 1

		currentId = 5 * state

	}
	fmt.Println("current Id", currentId)

	var count int64
	initializers.DB.Model(&post).Count(&count)

	result := initializers.DB.Where("ID > ? ", currentId).Limit(5).Find(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}

	totalPage := math.Ceil(float64(count) / float64(5))
	fmt.Println(totalPage)

	if len(posts) == 0 {
		c.JSON(404, gin.H{"status": 404, "message": "Bad Request", "Post": posts, "TotalPage": totalPage, "CurrentPage": "Null"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": posts, "TotalPage": totalPage, "CurrentPage": currentPage})
}
