package main

import (
	"fmt"
)

const prefix = "Hello, "

// Hello returns "Hello, world"
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
