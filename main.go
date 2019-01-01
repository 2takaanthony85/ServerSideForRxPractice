package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wakashiyo/serverSideForRxPractice/todos"
	"github.com/wakashiyo/serverSideForRxPractice/users"
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
		//users
		v1.GET("/user", users.SignIn)
		v1.POST("/user", users.SignUp)
		//v1.PUT("/user", users.SignOut)

		//todos
		v1.GET("/user/:id/todos", todos.GetTodos)
		v1.POST("/user/:id/todos", todos.CreateTodo)
		v1.PUT("/user/:id/todos/:todo-id", todos.UpdateTodo)
		v1.DELETE("/user/:id/todos/:todo-id", todos.DeleteTodo)
	}

	router.Run(":9000")
}
