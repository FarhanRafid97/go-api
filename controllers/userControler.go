package controllers

import (
	"go-api/helpers"
	"go-api/initializers"
	"go-api/models"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	user := models.User{}
	c.BindJSON(&user)

	if user.Password == "" || user.Username == "" {
		c.JSON(400, gin.H{"Status": 400, "Message": "please input username or password"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		errs := strings.Split(result.Error.Error(), "(")[0]
		c.JSON(400, gin.H{"Status": 400, "Message": strings.Replace(errs, `"`, "", 100)})
		return
	}

	c.JSON(200, gin.H{"Message": "Success create new User", "Data": user})
}

func LoginUser(c *gin.Context) {

	input := models.User{}
	c.BindJSON(&input)

	resultUser := models.User{}
	if input.Password == "" || input.Username == "" {
		c.JSON(400, gin.H{"Status": 400, "Message": "please input username or password"})
		return
	}

	result := initializers.DB.Where("username = ?", input.Username).First(&resultUser)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {

			c.JSON(404, gin.H{"Status": 404, "Message": "Username Or Password Invalid"})
			return
		}
		c.JSON(400, gin.H{"Status": 400, "Message": result.Error.Error()})
		return

	}
	err := helpers.ComparePassword(resultUser.Password, input.Password)
	if err != nil {
		panic(err)

	}
	accessTokenData := map[string]interface{}{"ID": resultUser.ID}
	token, err := helpers.Sign(accessTokenData, "ACCESS_TOKEN")
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"Message": "Success Login", "AceesToken": token})
}
