package main

import "fmt"

func main() {
	x := 5
	x = increment(x)

	fmt.Println(x)
}

func increment(x int) int {
	x++
	return x
}
