package main

import "fmt"

func main() {
	var n, f, x, y, z int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&f)
		sum := 0

		for j := 0; j < f; j++ {
			fmt.Scan(&x, &y, &z)
			sum += x * z
		}

		fmt.Println(sum)
	}
}
