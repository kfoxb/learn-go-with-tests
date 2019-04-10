package main

import (
	"fmt"
)

const prefixEnd = ", "
const englishPrefix = "Hello" + prefixEnd
const spanishPrefix = "Hola" + prefixEnd
const frenchPrefix = "Bonjour" + prefixEnd
const spanish = "Spanish"
const french = "French"

// Hello returns "Hello, world"
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
