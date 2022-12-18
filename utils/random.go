package utils

import "math/rand"

const (
	hex string = "01234567890abcdef"
)

func RandomHash(length int) string {
	arr := make([]byte, length)
	for i := range arr {
		index := rand.Intn(len(hex))
		arr[i] = hex[index]
	}
	return string(arr)
}
