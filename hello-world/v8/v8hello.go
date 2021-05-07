package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
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
	fmt.Println(Hello("lrg", "French"))
}
