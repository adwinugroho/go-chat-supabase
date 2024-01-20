package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(lengthNumber int) int {
	var charset = "0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, lengthNumber)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	byteToInt, _ := strconv.Atoi(string(b))
	return byteToInt
}
