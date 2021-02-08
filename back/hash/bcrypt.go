package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// BcryptPasswordHasher is an implementation of a users PasswordHasher that use Bcrypt.
type BcryptPasswordHasher struct{}

// Hash a password using bcrypt and returns the result.
func (hasher BcryptPasswordHasher) Hash(password string) string {
	rawResult, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(rawResult)
}

// CheckHashPassword compares users hashed password and a raw password and returns if are the same or not.
func (hasher BcryptPasswordHasher) CheckHashPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil
}
