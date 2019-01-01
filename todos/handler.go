package todos

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetTodos get user's todos by user-id
func GetTodos(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.String(422, "not covert type int")
		return
	}
	fmt.Printf("GetTodos id: %d", id)
	c.String(http.StatusOK, "can convert type int")
}

//CreateTodo create new todo
func CreateTodo(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		c.JSON(422, gin.H{
			"status":  "not posted",
			"message": "can not create todo",
			"nick":    "nick",
		})
		return
	}
	fmt.Printf("CreateTodo id: %d", id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": "can create todo",
		"nick":    "nick",
	})
}

//DeleteTodo delete user's todo by user-id and todo-id
func DeleteTodo(c *gin.Context) {
	fmt.Printf("Delete Todo")
	c.JSON(http.StatusOK, gin.H{
		"status":  "deleted",
		"message": "can delete todo",
		"nick":    "nick",
	})
}

//UpdateTodo update title or Contents of user's todo by user-id and toto-id
func UpdateTodo(c *gin.Context) {
	fmt.Printf("Update Todo")
	c.JSON(http.StatusOK, gin.H{
		"status":  "putted",
		"message": "can update todo",
		"nick":    "nick",
	})
}
