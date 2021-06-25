package pkg

import (
	"math/rand"
	"strings"
	"time"
)

func BreakStringByLine(str string) []string {
	return strings.Split(str, "\n")
}

func GetRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
