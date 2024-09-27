package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishPrefix      = "Hola, "
	frenchPrefix       = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Aaron", ""))
}
