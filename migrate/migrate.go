package main

import (
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}