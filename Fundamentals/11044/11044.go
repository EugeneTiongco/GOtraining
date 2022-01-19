package main

import "fmt"

func main() {
	var t, x, y int

	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&x, &y)
		sonars := (x / 3) * (y / 3)

		fmt.Println(sonars)
	}
}
