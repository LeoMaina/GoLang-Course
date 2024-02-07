package main

import "fmt"

type user struct {
	name        string
	phoneNumber int
}

type sender struct {
	rateLimit int
	user
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.phoneNumber)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("========================================")
}

func main() {
	test(sender{
		rateLimit: 1000,
		user: user{
			name:        "Leonard",
			phoneNumber: 254703663815,
		},
	})
}
