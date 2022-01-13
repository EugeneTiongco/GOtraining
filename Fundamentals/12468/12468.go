package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64

	n, _ := fmt.Scan(&x, &y)

	for n == 2 && x != -1 && y != -1 {
		if math.Abs(x-y) <= 50 {
			fmt.Println(math.Abs(x - y))
		} else if math.Abs(x-y) > 50 {
			fmt.Println(100 - Max(x, y) + Min(x, y))
		}
		n, _ = fmt.Scan(&x, &y)
	}
}

func Max(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Min(x float64, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
