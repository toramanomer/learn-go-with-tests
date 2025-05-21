package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishDefaultRecipient = "World"

func Hello(recipient string) string {
	if recipient == "" {
		recipient = englishDefaultRecipient
	}
	return englishHelloPrefix + recipient
}

func main() {
	fmt.Println(Hello("World"))
}
