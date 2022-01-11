package main

import "fmt"

func main() {
	var t, x, y int

	fmt.Scan(&t)
	if t < 15 {
		n, _ := fmt.Scan(&x, &y)
		for n == 2 && t > 0 {
			switch {
			case x > y:
				fmt.Println(">")
			case x < y:
				fmt.Println("<")
			case x == y:
				fmt.Println("=")
			}
			t--
			if t > 0 {
				n, _ = fmt.Scan(&x, &y)
			}
		}
	}
}
