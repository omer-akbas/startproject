package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	fmt.Println("Hello World.")
	fmt.Println(randToken(3))
	fmt.Printf("%8s %8s\n", "Name", "Password")
	fmt.Printf("%8s %8s\n", "-----", "--------")
	fmt.Printf("%8s %8s\n", "Foo", randToken(3))
	fmt.Printf("%8s %8s\n", "Bar", randToken(3))
	fmt.Printf("%8s %8s\n", "Baz", randToken(3))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
