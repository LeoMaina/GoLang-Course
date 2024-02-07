package main

import "fmt"

func main() {

	//One way of writing the if statements in Go (norm)
	length := getLength(email)
	if length > 1 {
		fmt.Println("Email is valid")
	}

	//The alternative way of doing this which keeps the code concise and scope limited
	if length := getLength(email); length > 1 {
		fmt.Println("Email is valid")
	}
}
