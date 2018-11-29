package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("Hello World.")
	fmt.Println(randToken(3))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
