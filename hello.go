package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	englishIntroPrefix = "Hello, "
	spanishIntroPrefix = "Hola, "
	frenchIntroPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("jun", ""))
}

func getGreetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishIntroPrefix
	case french:
		prefix = frenchIntroPrefix
	default:
		prefix = englishIntroPrefix
	}

	return
}
