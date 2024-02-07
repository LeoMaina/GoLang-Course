package main

import "fmt"

//Defining a struct
type messageToSend struct {
	phoneNumber int
	message     string
}

func test(var1 messageToSend) {
	fmt.Printf("Sending this message: %s to this contact: %v\n", var1.message, var1.phoneNumber)
	fmt.Println("====================================")

}

func main() {
	test(messageToSend{
		phoneNumber: 254703663815,
		message:     "You have a bright future bro!",
	})
	test(messageToSend{
		phoneNumber: 254712780646,
		message:     "What's popping mzee?",
	})
	test(messageToSend{
		phoneNumber: 254726022485,
		message:     "Umekachora aje?",
	})

}
