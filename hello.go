package main

import (
	"fmt"
)

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola! "
	frenchHelloPrefix  = "Bonjour! "
)

func Hello(name, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"

	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("Ashmit", ""))
}
