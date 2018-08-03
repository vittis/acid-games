package iogames

import (
	"math/rand"
	"time"
)

//Contains .
func Contains(arr [5]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//RandomStringWithCharset .
func RandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//RandomString .
func RandomString(length int) string {
	return RandomStringWithCharset(length, charset)
}
