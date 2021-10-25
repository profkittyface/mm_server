package main

import (
  "crypto/sha256"
  "crypto/sha1"
  "encoding/hex"
  "time"
  "context"
)

func GenerateLocationKey(username string) string {
  s := "mysterymeeting-" + username
  hasher := sha1.New()
  bytearray := []byte(s)
  hasher.Write(bytearray)
  byteres := hasher.Sum(nil)
  res := hex.EncodeToString(byteres)
  return res
}

func HashPassword(password string) string {
  hasher := sha256.New()
  bytearray := []byte(password)
  hasher.Write(bytearray)
  byteres := hasher.Sum(nil)
  res := hex.EncodeToString(byteres)
  return res
}

func GenerateLoginKey(username string) string {
  t := time.Now()
  s := username + "-" + t.String()
  hasher := sha1.New()
  bytearray := []byte(s)
  hasher.Write(bytearray)
  byteres := hasher.Sum(nil)
  res := hex.EncodeToString(byteres)
  return res
}

func CheckUsernameAndPassword(username string, password string) bool {
  u := User{}
  ctx := context.Background()
  db := getCursor()
  db.NewSelect().Model(&u).Where("username = ?", username).Scan(ctx)
  if u.Id == 0 {
    return false
  }
  if u.Password == HashPassword(password){
    return true
  }
  return false
}

func CheckCookieKey(cookie_key string) bool {
	ctx := context.Background()
	db := getCursor()
	lt := LoginTrack{}
	db.NewSelect().Model(&lt).Where("cookie_key = ?", cookie_key).Scan(ctx)
	if lt.Id == 0 {
		return false
	}
  if time.Now().After(lt.Expires) {
    db.NewDelete().Model(&lt).Where("cookie_key = ?", cookie_key).Exec(ctx)
    return false
  }
  return true
}
