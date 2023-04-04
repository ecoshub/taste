package utils

import (
	"math/rand"
	"time"
)

const (
	hexChars string = "01234567890abcdef"
)

// RandomHash generates a random hexadecimal string of the given length.
// It seeds the random number generator with the current time, and uses it
// to pick a random character from the set of hexadecimal digits to build
// the random hash.
func RandomHash(length int) string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate an array of random bytes
	hash := make([]byte, length)
	for i := range hash {
		// Pick a random character from the set of hexadecimal digits.
		index := rand.Intn(len(hexChars))
		hash[i] = hexChars[index]
	}

	// Convert the byte array to a string and return it
	return string(hash)
}
