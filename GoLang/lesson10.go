package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {

	fmt.Println(concat("And I", " Am Iron Man"))
	fmt.Println(concat("You will", " bounce back from the heartbreak"))
}
