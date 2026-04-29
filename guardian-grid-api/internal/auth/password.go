// internal/auth/password.go
package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(hash)
}

func CheckPassword(input, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input)) == nil
}
