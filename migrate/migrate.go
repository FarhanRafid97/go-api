package main

import (
	"go-api/initializers"
	"go-api/models"
)

func init() {
	initializers.ConnectToDb()
	initializers.InitializerEnvVariable()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
