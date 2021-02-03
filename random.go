package utils

import "math/rand"


func RandomString(n int) string {
	return RandomStringRange(n, Chars)
}

func RandomStringRange(n int, chars string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

