package main

import "fmt"

func main() {
	var username string = "User1"
	var password string = "36529746"

	//Print
	fmt.Println("Authorization: Basic", username+":"+password)
}
