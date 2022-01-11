package main

import "fmt"

func main() {
	var x, y int

	n, _ := fmt.Scan(&x, &y)
	for n == 2 {
		fmt.Println(y - x)
		n, _ = fmt.Scan(&x, &y)
	}
}
