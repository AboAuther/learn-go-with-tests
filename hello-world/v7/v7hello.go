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
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}
func main() {
	fmt.Println(Hello("lrg", "French"))
}
