package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

func generate(charset string, length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}

func checksum(id string) string {
	var sum int
	for i := 0; i < 12; i++ {
		n, _ := strconv.Atoi(string(id[i]))
		sum += n * (13 - i)
	}
	r := 11 - (sum % 11)
	if r > 9 {
		r -= 10
	}
	return strconv.Itoa(r)
}

func main() {
	id := generate("12345678", 1) + generate("0123456789", 11)
	id = id + checksum(id)
	fmt.Println(id)
}
