package hashing

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 computes the MD5 hash of the input string and returns it as a hexadecimal string.
// It uses the crypto/md5 package to create a new MD5 hash, writes the input string to the hash,
// and then converts the resulting byte slice to a hexadecimal string.
func MD5(input string) string {
	// create a new MD5 hash
	hash := md5.New()

	// write the input string to the hash
	hash.Write([]byte(input))

	// get the resulting hash as a byte slice
	hashBytes := hash.Sum(nil)

	// convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}

// SHA256 computes the SHA256 hash of the input string and returns it as a hexadecimal string.
// It uses the crypto/sha256 package to create a new SHA256 hash, writes the input string to the hash,
// and then converts the resulting byte slice to a hexadecimal string.
func SHA256(input string) string {
	// create a new SHA256 hash
	hash := sha256.New()

	// write the input string to the hash
	hash.Write([]byte(input))

	// get the resulting hash as a byte slice
	hashBytes := hash.Sum(nil)

	// convert the byte slice to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}
