package main

import "fmt"

func main() {
	var C, N, x int

	fmt.Scan(&C)

	for i := 0; i < C; i++ {
		fmt.Scan(&N)
		students := make([]int, N)
		total := 0
		count := 0

		for j := 0; j < N; j++ {
			fmt.Scan(&x)
			students[j] = x
			total += x
		}
		average := float64(total) / float64(N)

		for j := 0; j < N; j++ {
			if float64(students[j]) > average {
				count++
			}
		}
		fmt.Printf("%.3f%%\n", float64(count)/float64(N)*100)
	}
}
