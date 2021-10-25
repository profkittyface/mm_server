package main

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"users"`
	Id            int       `bun:"id"`
	Username      string    `bun:"username"`
	Password      string    `bun:"password"`
	Email         string    `bun:"email"`
	FirstName     string    `bun:"first_name"`
	LastName      string    `bun:"last_name"`
	LastIp        string    `bun:"lastip"`
	LastLogin     string    `bun:"last_login"`
	DateCreated   time.Time `bun:"date_created,default:current_timestamp"`
}

type Location struct {
	bun.BaseModel `bun:"location"`
	Id            int       `bun:"id"`
	Location      string    `bun:"location"`
	UserId        int       `bun:"userid"`
	LastUpdated   time.Time `bun:"last_updated,default:current_timestamp"`
}

type Event struct {
	bun.BaseModel `bun:"event"`
	Id            int       `bun:"id"`
	Name          string    `bun:"name"`
	Venue         string    `bun:"venue"`
	MeetingTime   time.Time `bun:"meeting_time"`
	Attended      string    `bun:"attended"`
	DateCreated   time.Time `bun:"date_created,default:current_timestamp"`
	LastUpdated   time.Time `bun:"last_updated,default:current_timestamp"`
}

type Venue struct {
	bun.BaseModel `bun:"venue"`
	Id            int       `bun:"id"`
	Name          string    `bun:"name"`
	Address       string    `bun:"address"`
	Hours         string    `bun:"hours"`
	DateCreated   time.Time `bun:"date_created,default:current_timestamp"`
	LastUpdated   time.Time `bun:"last_updated,default:current_timestamp"`
}

type LocationKey struct {
	bun.BaseModel `bun:"location_key"`
	Id            int       `bun:"id"`
	Userid        int       `bun:"userid"`
	Key           string    `bun:"key"`
	DateCreated   time.Time `bun:"date_created,default:current_timestamp"`
	LastUpdated   time.Time `bun:"last_updated,default:current_timestamp"`
}

type LoginTrack struct {
	bun.BaseModel `bun:"auth_key"`
	Id            int       `bun:"id"`
	Userid        int       `bun:"userid"`
	CookieKey     string    `bun:"cookie_key"`
	Expires       time.Time `bun:"expires"`
}

type Profile struct {
	bun.BaseModel `bun:"profile"`
	Id            int    `bun:"id"`
	Userid        int    `bun:"userid"`
	AboutMe       string `bun:"about_me"`
	Interests     string `bun:"interests"`
	Location      string `bun:"location"`
}

func getUserFromUsername(username string) (User, error) {
	ctx := context.Background()
	db := getCursor()
	user := User{}
	db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	if user.Id == 0 {
		fmt.Println("Username not found")
	}
	return user, nil
}

func getUserFromId(userid int) User {
	ctx := context.Background()
	db := getCursor()
	user := User{}
	db.NewSelect().Model(&user).Where("id = ?", userid).Scan(ctx)
	if user.Id == 0 {
		fmt.Println("Userid not found")
		return user
	}
	return user
}

func getIdFromUser(username string) int {
	ctx := context.Background()
	db := getCursor()
	user := User{}
	db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	return user.Id
}

func getUserFromCookieKey(cookie_key string) User {
  user := User{}
  ctx := context.Background()
  db := getCursor()
  lt := LoginTrack{}
  db.NewSelect().Model(&lt).Where("cookie_key = ?", cookie_key).Scan(ctx)
  db.NewSelect().Model(&user).Where("userid = ?", lt.Userid).Scan(ctx)
  return user
}

func getProfileFromId(userid int) Profile {
  profile := Profile{}
  ctx := context.Background()
  db := getCursor()
  db.NewSelect().Model(&profile).Where("userid = ?", userid).Scan(ctx)
  return profile
}
