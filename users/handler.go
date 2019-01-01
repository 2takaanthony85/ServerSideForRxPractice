package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SignIn signin user
func SignIn(c *gin.Context) {
	mail := c.Query("email")
	pass := c.Query("password")
	fmt.Printf("signin --- email: %s, pass: %s", mail, pass)
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": "sign in",
		"nick":    "nick",
	})
}

//SignUp signup user
func SignUp(c *gin.Context) {
	mail := c.Query("email")
	pass := c.Query("password")
	fmt.Printf("signup --- email: %s, pass: %s", mail, pass)
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": "sign up",
		"nick":    "nick",
	})
}

//SignOut signout user
func SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": "sign out",
		"nick":    "nick",
	})
}
