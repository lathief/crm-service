package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 8

	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), salt)
	return string(hash)
}

func ComparePass(hash, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
