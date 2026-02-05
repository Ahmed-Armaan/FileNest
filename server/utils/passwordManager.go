package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashedPassword string, givenPassword string) bool {
	hashByte := []byte(hashedPassword)
	givenByte := []byte(givenPassword)
	if err := bcrypt.CompareHashAndPassword(hashByte, givenByte); err != nil {
		return false
	}
	return true
}
