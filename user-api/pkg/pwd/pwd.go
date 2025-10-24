package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash generates a bcrypt hash of the password
func Hash(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// Match compares a hashed password with a plain password
func Match(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
