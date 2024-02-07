package main

import "fmt"

func main() {
	firstName, _ := getName()

	fmt.Println("Welcome to TextIo", firstName)
}

func getName() (string, string) {
	return "John", "Doe"
}
