package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidatePasswordHash(password, hashedPassword string) bool {
	fmt.Println("entering here")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	print(err)
	return err == nil
}
