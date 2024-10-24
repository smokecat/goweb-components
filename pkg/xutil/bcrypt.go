package xutil

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultBcryptCost = 11
)

func BcryptGeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultBcryptCost)
	if err != nil {
		return "", errors.New("Failed to generate password hash")
	}
	return string(bytes), nil
}

func BcryptCheckPassword(hashed, password string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
