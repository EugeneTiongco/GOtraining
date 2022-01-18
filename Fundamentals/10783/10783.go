package main

import (
	"fmt"
)

func main() {
	var T, a, b int
	count := 1
	fmt.Scan(&T)

	if T >= 1 && T <= 100 {
		_, err := fmt.Scan(&a)
		_, err = fmt.Scan(&b)
		for err == nil && count <= T {
			if a%2 == 0 {
				fmt.Printf("Case %v: %v\n", count, getSum(a+1, b))
			} else {
				fmt.Printf("Case %v: %v\n", count, getSum(a, b))
			}
			count++
			if count <= T {
				_, err = fmt.Scan(&a)
				_, err = fmt.Scan(&b)
			}
		}
	}
}

func getSum(a int, b int) int {
	sum := a
	count := (b - a) / 2

	for count != 0 {
		sum = sum + a + power(2, count)
		count--
	}

	return sum
}

func power(a int, b int) int {
	pow := 1
	for i := 0; i < b; i++ {
		pow = pow * a
	}
	return pow
}
