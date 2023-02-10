package security

import "golang.org/x/crypto/bcrypt"

// HashMake is function to hash
func HashMake(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

// HashVerify is function to verify hashed password
func HashVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
