package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Hola, "
const spanishHelloPrefix = "Hallo, "

func Hello(name, language string) string {
	prefix := englishHelloPrefix
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Jahn", "english"))
}
