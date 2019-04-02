package main

import (
	"fmt"
)

const prefix = "Hello, "

// Hello returns "Hello, world"
func Hello(name string) string {
	return prefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
