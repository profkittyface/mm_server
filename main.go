package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"context"
)

func main() {
	router := gin.Default()

	router.GET("/", dashboard)

	router.GET("/login", login)
	router.POST("/login", postLogin)
	router.GET("/logout", logout)

	router.GET("/debug", debug)
	router.GET("/db", db)

	router.Run(":8080")
}

func debug(c *gin.Context){
	c.String(http.StatusOK, HashPassword("ahunt"))
}



func login(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	c.SetCookie("mm_auth", "test", 3600, "/", "localhost", false, true)
	c.SetCookie("mm_user", username, 3600, "/", "localhost", false, true)
	cookie, _ := c.Cookie("mm_login")
	fmt.Println(cookie)
}

func postLogin(c *gin.Context){
	c.SetCookie("mm_login", "test", 3600, "/", "localhost", false, true)
	cookie, _ := c.Cookie("mm_login")
	fmt.Println(cookie)
}

func logout(c *gin.Context){
	// return true
}

func dashboard(c *gin.Context){
	// return true
}
