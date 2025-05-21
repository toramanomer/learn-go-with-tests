package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(recipient string) string {
	return englishHelloPrefix + recipient
}

func main() {
	fmt.Println(Hello("World"))
}
