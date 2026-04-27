package crypto

import (
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	// Hash password with Sha512 (resistant to modern attacks after 2030) to prevent DoS attacks via Argon2
	sha := sha512.New()
	sha.Write([]byte(password))
	passwordHash := sha.Sum(nil)

	// Generate salt
	salt := GenerateRandomBytes(20)

	// Hash Password
	passwordHash = argon2.IDKey(passwordHash, salt, 2, 19*1024, 1, 32)

	// Encode salt and hash and return the two
	return base64.StdEncoding.EncodeToString(passwordHash) + "." + base64.StdEncoding.EncodeToString(salt), nil
}

func VerifyPassword(password, hash string) (bool, error) {
	// Unwrap current hash and salt
	s := strings.Split(hash, ".")
	if len(s) != 2 {
		return false, errors.New("invalid hash format")
	}

	// Decode password hash
	passwordHash, err := base64.StdEncoding.DecodeString(s[0])
	if err != nil {
		return false, errors.New("invalid hash format")
	}

	// Decode salt
	salt, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false, errors.New("invalid hash format")
	}

	// Hash new password with Sha512
	sha := sha512.New()
	sha.Write([]byte(password))
	newHash := sha.Sum(nil)

	// Hash new password with Argon2ID
	newHash = argon2.IDKey(newHash, salt, 2, 19*1024, 1, 32)

	// Prevent timing attacks
	return subtle.ConstantTimeCompare(newHash, passwordHash) == 1, nil
}
