package main

import (
	"github.com/betology/go-crud/initializers"
	"github.com/betology/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
