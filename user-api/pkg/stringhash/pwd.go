package stringhash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash generates a bcrypt hash of the string
func Hash(stringToHash string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(stringToHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Match compares a hashed string with a plain string
func Match(hashedString, plainString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(plainString))
	return err == nil
}
