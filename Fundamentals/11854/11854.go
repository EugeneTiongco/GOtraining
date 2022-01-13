package main

import "fmt"

func main() {
	var x, y, z int

	n, _ := fmt.Scan(&x, &y, &z)
	for n == 3 && x != 0 && y != 0 && z != 0 {
		if x*x+y*y == z*z {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
		n, _ = fmt.Scan(&x, &y, &z)
	}
}
