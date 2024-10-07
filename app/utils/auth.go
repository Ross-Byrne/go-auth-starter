// Authentication and Authorisation related utils

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const DEFAULT_HASH_COST int = 14

func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DEFAULT_HASH_COST)
	return string(bytes), err
}

func CompareHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
