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
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
