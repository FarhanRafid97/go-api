package controllers

import (
	"fmt"
	"go-api/initializers"
	"go-api/models"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IdPost struct {
	Id int
}

func CreatePost(c *gin.Context) {

	post := models.Post{}
	c.BindJSON(&post)
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": post})
}

func GetPost(c *gin.Context) {
	params := c.Params
	id, _ := params.Get("id")

	post := models.Post{}
	result := initializers.DB.Where("ID = ?", id).First(&post)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {

			c.JSON(404, gin.H{"Status": 404, "Message": result.Error.Error()})
			return
		}
		c.JSON(400, gin.H{"Status": 400, "Message": result.Error.Error()})
		return

	}

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": post})
}

func GetPosts(c *gin.Context) {
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
	param := c.Request.URL.Query()
	page := 0
	limit := 5
	paramPage := param.Get("page")
	paramLimit := param.Get("limit")
	if paramLimit != "" {

		limitNum, _ := strconv.Atoi(paramLimit)
		limit = limitNum
	}
	if paramPage != "" {

		pageNum, _ := strconv.Atoi(paramPage)
		page = pageNum
	}
	var count int64
	result2nd := initializers.DB.Model(&post).Count(&count)
	if result2nd.Error != nil {
		c.Status(401)
		return
	}
	totalPage := math.Ceil(float64(count) / float64(limit))

	result := initializers.DB.Limit(limit).Offset(page * limit).Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	if len(posts) == 0 {
		c.JSON(404, gin.H{"status": 404, "message": "Bad Request", "Post": posts, "TotalPage": totalPage, "CurrentPage": "Null"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "OK", "Post": posts, "TotalPage": totalPage})
}

func DeletePost(c *gin.Context) {
	post := models.Post{}
	params := c.Params
	id, _ := params.Get("id")

	numId, _ := strconv.Atoi(id)
	isPost := initializers.DB.Where("id = ?", numId).First(&post)

	if isPost.Error != nil {
		if isPost.Error.Error() == "record not found" {
			c.JSON(404, gin.H{"Message": isPost.Error.Error(), "Status": 404})
		}
		c.JSON(400, gin.H{"Message": isPost.Error.Error(), "Status": 400})

		return
	}

	result := initializers.DB.Where("id = ?", numId).Delete(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "Success Delete Data"})
}

func UpdatePost(c *gin.Context) {
	params := c.Params
	id, _ := params.Get("id")
	post := models.Post{}

	numId, _ := strconv.Atoi(id)
	isPost := initializers.DB.Where("id = ?", numId).First(&post)

	if isPost.Error != nil {
		if isPost.Error.Error() == "record not found" {
			c.JSON(404, gin.H{"Message": isPost.Error.Error(), "Status": 404})
		}
		c.JSON(400, gin.H{"Message": isPost.Error.Error(), "Status": 400})

		return
	}
	c.BindJSON(&post)
	result := initializers.DB.Where("id = ?", numId).Updates(&models.Post{Title: post.Title, Body: post.Body}).Take(&post)
	if result.Error != nil {
		c.JSON(400, gin.H{"Message": result.Error.Error(), "Status": 400})
		return
	}

	c.JSON(200, gin.H{"status": 200, "message": "Success Delete Data", "Data": post})
}
