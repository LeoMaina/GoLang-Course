package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}
type user struct {
	name        string
	phoneNumber int
}

func canSendMessage(var1 messageToSend) bool {
	if var1.message == "" || var1.sender.name == "" || var1.sender.phoneNumber == 0 || var1.recipient.name == "" || var1.recipient.phoneNumber == 0 {
		return false
	} else {
		return true
	}
}

func test(var1 messageToSend) {
	fmt.Printf(`Sending "%s" to %s (%v) from %s (%v) ...`,
		var1.message,
		var1.recipient.name,
		var1.recipient.phoneNumber,
		var1.sender.name,
		var1.sender.phoneNumber)

	if canSendMessage(var1) {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Cannot send message")
	}
	fmt.Println("==========================================")
}

func main() {
	test(messageToSend{
		message: "Hello, I am Go Developer",
		sender: user{
			name:        "Leonard",
			phoneNumber: 254703663815,
		},
		recipient: user{
			name:        "The_World",
			phoneNumber: 2547000000,
		},
	})

	test(messageToSend{
		message: "What's up?",
		sender: user{
			name:        "Phoebe",
			phoneNumber: 254718905769,
		},
		recipient: user{
			name:        "Bazenga",
			phoneNumber: 254712780646,
		},
	})
}
