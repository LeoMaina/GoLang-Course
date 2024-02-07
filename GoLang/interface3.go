/*
Multiple Interfaces
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.

Assignment:
Add the required methods so that the email type implements both the expense and printer interfaces.

cost()
If the email is not "subscribed", then the cost is 0.05 for each character in the body. If it is, then the cost is 0.01 per character.

print()
The print method should print to standard out the email's body text.


Solution: Define the cost method to receive the type email and to return a float64 value when called. We change the length of the e.body to float64 since you cannot mutliply
			an int and a float value in Go. Then you define the print method to receive the type email and to print the body of the type email to the standard output.
*/

package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if e.isSubscribed {
		return 0.01 * float64(len(e.body))
	} else {
		return 0.05 * float64(len(e.body))
	}
}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: %.2f...\n", e.cost())
	p.print()
	fmt.Println("==================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "Hello there,",
	}
	test(e, e)

	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
}
