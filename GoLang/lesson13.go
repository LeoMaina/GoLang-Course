package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return //This is an implicit return, an explicit return would take the form "return(var1, var2, var3)"
}

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)

	if yearsUntilAdult == 0 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are an adult in", yearsUntilAdult, "years")
	}
	if yearsUntilDrinking == 0 {
		fmt.Println("You can drink")
	} else {
		fmt.Println("You can drink in", yearsUntilDrinking, "years")
	}
	if yearsUntilCarRental == 0 {
		fmt.Println("You can rent a car")
	} else {
		fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	}
}
func main() {
	test(8)
	test(16)
	test(24)
	test(32)
}
