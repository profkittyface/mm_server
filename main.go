package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	"context"
	"time"
)

func main() {
	router := gin.Default()
	// router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/login", getLogin)
	router.POST("/login", postLogin)
	router.GET("/logout", getLogout)

	router.GET("/", getDashboard)
	router.GET("/events", getEvent)
	router.GET("/profile", getProfile)

	router.GET("/debug", debug)

	router.Run(":8080")
}

func debug(c *gin.Context){
	GetDailyCluster()
}

type Dashboard struct {
	User User
	Profile Profile
}

type NewEvent struct {
	Title string
	Description string

}

func getDashboard(c *gin.Context){
	// User object
	// Profile object

	auth, _ := c.Cookie("mm_auth")
	if CheckCookieKey(auth){
		user := getUserFromCookieKey(auth)
		profile := getProfileFromId(user.Id)
		dashboard := Dashboard{User: user, Profile: profile}
		c.JSON(200, dashboard)
	}
	c.JSON(200, gin.H{"message": "not dashboard"})
}

func getEvent(c *gin.Context){
	//
}

func getProfile(c *gin.Context){
	//
}

func getLogin(c *gin.Context){
	auth, _ := c.Cookie("mm_auth")
	someSortOfAuthObject := map[string]interface{} {"auth": auth}
	c.HTML(http.StatusOK, "login.tmpl", someSortOfAuthObject)
}

func postLogin(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	if CheckUsernameAndPassword(username, password){
		authKey := GenerateLoginKey(username)
		c.SetCookie("mm_auth", authKey, 3600, "/", "localhost", false, true)
		c.SetCookie("mm_user", username, 3600, "/", "localhost", false, true)
		lt := LoginTrack{}
		lt.Userid = getIdFromUser(username)
		lt.CookieKey = authKey
		lt.Expires = time.Now().Add(time.Hour * 1)
		ctx := context.Background()
		db := getCursor()
	  db.NewInsert().Model(&lt).Exec(ctx)
		c.JSON(200, gin.H{"message": "logged_in"})
	}
}

func getLogout(c *gin.Context){
	c.SetCookie("mm_auth", "", 3600, "/", "localhost", false, true)
}
