package main

import (
	"math/rand"
	"time"
)

var (
	words []string
)

const (
	N = 1024
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func init() {
	w := make([]string, N)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < N; i++ {
		l := r.Int() % N
		w[i] = GetRandomString(l)
	}
	words = w
}

func main() {

}
