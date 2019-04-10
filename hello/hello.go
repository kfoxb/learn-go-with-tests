package main

import (
	"fmt"
)

const prefixEnd = ", "
const englishPrefix = "Hello" + prefixEnd
const spanishPrefix = "Hola" + prefixEnd
const spanish = "Spanish"

// Hello returns "Hello, world"
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefix
	if language == spanish {
		prefix = spanishPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
