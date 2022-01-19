package main

import "fmt"

const one = "one"

func main() {
	var t int
	var word string

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&word)
		fmt.Println(check(word))
	}
}

func check(word string) int {
	var diff int
	if len(word) > 3 {
		return 3
	}

	for i := 0; i < 3; i++ {
		if word[i] != one[i] {
			diff++
		}
	}

	if diff == 1 {
		return 1
	}

	return 2
}
