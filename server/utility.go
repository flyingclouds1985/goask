package server

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func checkNil(item interface{}, message string) {
	var err string
	if e, ok := item.(error); ok {
		err = e.Error()
	}
	if item != nil {
		fmt.Printf("Error: %s, Message %s", err, message)
	}
}

// HashPassword makes password in bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash checks hashed password against string password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RandomString generates random string in desire length.
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
