package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/todos")
	{
		v1.POST("/", createTodo)
		//v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/api/todos", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "dfdfd",
		})
	})
	router.Run(":8000")
}

func createTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hohoho",
	})
}

func fetchAllTodo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "addAll"})
}

func fetchSingleTodo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "addSingle"})
}

func updateTodo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "addupdate"})
}

func deleteTodo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "del"})
}
