package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphanums = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt return an int64 between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString return an string of length n
func RandomString(n int) string {
	var sb strings.Builder
	var len = len(alphanums)
	var c byte
	for i := 0; i < n; i++ {
		c = alphanums[rand.Intn(len)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName return a random name
func RandomName() string {
	name := RandomString(7)
	name += " " + RandomString(9)
	return name
}

// RandomUsername returns a 32 bit API key
func RandomUsername() string {
	return RandomString(12)
}

// RandomAPIKey returns a 32 bit API key
func RandomAPIKey() string {
	return RandomString(32)
}

// RandomKey returns a 32 bit API key
func RandomKey() string {
	return RandomString(12)
}

// RandomValue returns a 32 bit API key
func RandomValue() string {
	return RandomString(50)
}
