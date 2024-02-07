package main

import (
	"fmt"
)

type email struct {
	isSubscribed bool
	toAddress    string
	body         string
}

type sms struct {
	isSubscribed  bool
	toPhoneNumber string
	body          string
}

type invalid struct{}

type expense interface {
	cost() float64
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	} else {
		return float64(len(e.body)) * 0.05
	}
}

func (s sms) cost() float64 {
	if s.isSubscribed {
		return float64(len(s.body)) * 0.03
	} else {
		return float64(len(s.body)) * 1
	}
}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	//This is how you use type-switches
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}
func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}
