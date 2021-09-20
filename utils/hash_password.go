package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword create hash of password
func HashPassword(password string) (*string, error) {
	rawHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hash := string(rawHash)
	return &hash, nil
}

// ComparePassword compare hash with password
func ComparePassword(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}
