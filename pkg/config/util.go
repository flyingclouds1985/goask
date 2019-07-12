package config

import (
	"log"
	"math/rand"
	"os"
)

// RandomString generates random string in desire length.
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}

// RouterSecretKey returns the key.
func routerSecretKey(len int) string {
	k := os.Getenv("RouterSecretKey")
	if k == "" {
		s := randomString(len)
		err := os.Setenv("RouterSecretKey", s)
		if err != nil {
			log.Println(err)
		}
		return s
	}
	return k
}
