package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {

	bytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", hashedPassword), nil
}

func ComparePassword(rawPassword, hashPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	if err != nil {
		return false
	} else {
		return true
	}
}
