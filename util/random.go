package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomText() string {
	return RandomString(6)
}
func RandomNum() int64 {
	return RandomInt(1, 15)
}
func RandomWeight() int64 {
	return RandomInt(100, 300)
}
func RandomHeight() int64 {
	return RandomInt(100, 250)
}
func RandomCal() int64 {
	return RandomInt(100, 800)
}
func RandomQuantity() int64 {
	return RandomInt(30, 100)
}
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
func RandomDate() string {
	startDate := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	duration := endDate.Sub(startDate)
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	randomDate := startDate.Add(randomDuration)
	return randomDate.Format("2006-01-02")
}
