package main

import "fmt"

func main() {
	var N, T int
	count := 1

	fmt.Scan(&T)

	if T <= 50 && T > 0 {
		n, _ := fmt.Scan(&N)
		for n == 1 && count <= T {
			var monsters [100]int
			for i := 0; i < N; i++ {
				fmt.Scan(&monsters[i])
			}
			for j := 1; j < N; j++ {
				if monsters[0] < monsters[j] {
					monsters[0] = monsters[j]
				}
			}
			fmt.Printf("Case %v: %v\n", count, monsters[0])
			count++
			if count <= T {
				n, _ = fmt.Scan(&N)
			}
		}
	} else {
		fmt.Println("Error: Input a positive integer less than or equal to 50")
	}

}
