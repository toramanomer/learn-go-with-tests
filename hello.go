package main

import "fmt"

func Hello(recipient string) string {
	return fmt.Sprintf("Hello, %s", recipient)
}

func main() {
	fmt.Println(Hello("World"))
}
