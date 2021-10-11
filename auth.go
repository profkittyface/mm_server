package main

import (
  "crypto/sha256"
  "crypto/sha1"
  "encoding/hex"
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
