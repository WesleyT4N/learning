package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const engHelloPrefix = "Hello, "
const spanHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) string {
    switch language {
        case spanish:
            return spanHelloPrefix
        case french:
            return frenchHelloPrefix
        default:
            return engHelloPrefix
    }
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
