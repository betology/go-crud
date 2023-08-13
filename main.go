package main

import (
	"github.com/betology/go-crud/controllers"
	"github.com/betology/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.Run()
}
