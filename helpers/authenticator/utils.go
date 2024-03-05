package authenticator

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
)

// checkPassword checks if the provided password matches the stored hashed password
func CheckPassword(storedHash string, plaintextPassword string) bool {
	// Split the stored hash into its components
	parts := strings.Split(storedHash, "$")
	iterations, _ := strconv.Atoi(parts[1])
	salt := []byte(parts[2])
	hashedPassword, _ := base64.StdEncoding.DecodeString(parts[3])

	// Derive a key using PBKDF2
	derivedKey := pbkdf2.Key([]byte(plaintextPassword), salt, iterations, len(hashedPassword), sha256.New)

	// Compare the derived key with the stored hashed password
	return subtle.ConstantTimeCompare(derivedKey, hashedPassword) == 1
}
