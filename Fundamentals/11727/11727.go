package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, x, y, z int
	count := 1

	fmt.Scan(&T)
	if T < 20 && T > 0 {
		n, _ := fmt.Scan(&x, &y, &z)
		for n == 3 && count <= T {
			employees := []int{x, y, z}
			sort.Ints(employees)
			fmt.Printf("Case %v: %v\n", count, employees[1])
			count++
			if count <= T {
				n, _ = fmt.Scan(&x, &y, &z)
			}
		}
	} else {
		fmt.Println("Error: Input a positive integer less than 20")
	}
}
