package main

import (
	"net/http"

	"github.com/2takaanthony85/serverSideForRxPractice/todos"
	"github.com/gin-gonic/gin"
)

func getHello() string {
	return "hello world!!"
}

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello world!!")
}

func main() {
	router := gin.Default()

	router.GET("/", helloworld)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/:id/todos", todos.GetTodos)
		v1.POST("/user/:id/todos", todos.CreateTodo)
		v1.PUT("/user/:id/todos/:todo-id", todos.UpdateTodo)
		v1.DELETE("/user/:id/todos/:todo-id", todos.DeleteTodo)
	}

	router.Run(":9000")
}
