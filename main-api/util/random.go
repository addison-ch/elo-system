package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random int between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// random string of length n
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]rune, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func RandomRating() int32 {
	return int32(RandomInt(500, 3000))
}

func RandomMatches() int32 {
	return int32(RandomInt(0, 100))
}

func RandomDescription() string {
	return RandomString(20)
}

func RandomUrl() string {
	collection := []string{
			"https://images.unsplash.com/photo-1529778873920-4da4926a72c2",
			"https://images.unsplash.com/photo-1615497001839-b0a0eac3274c",
			"https://images.unsplash.com/photo-1598935888738-cd2622bcd437",
	}

	n := len(collection)

	return collection[rand.Intn(n)]
}