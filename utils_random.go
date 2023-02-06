package taste

import "math/rand"

const (
	hex string = "01234567890abcdef"
)

// RandomHash generate random hash with given length
// generated hash is lowercase
func RandomHash(length int) string {
	arr := make([]byte, length)
	for i := range arr {
		index := rand.Intn(len(hex))
		arr[i] = hex[index]
	}
	return string(arr)
}
