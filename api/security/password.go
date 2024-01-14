package security

import "golang.org/x/crypto/bcrypt"

// Hash takes a string and returns a byte array and an error
// The byte array is the hashed version of the string
// The error is nil if the hashing was successful
func Hash(password string) ([]byte, error) {
 // GenerateFromPassword takes a byte array and a cost factor
 // The cost factor is the number of rounds used to generate the hashed version
 // The higher the cost factor, the more secure the hashed version is
 // DefaultCost is the recommended cost factor
 return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword takes a hashed password and a plain text password
// It returns an error if the plain text password does not match the hashed password
func VerifyPassword(hashedPassword, password string) error {
 // CompareHashAndPassword takes two byte arrays and compares them
 // It returns an error if the two byte arrays do not match
 return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
