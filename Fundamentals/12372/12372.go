package main

import "fmt"

func main() {
	var T, l, w, h int

	count := 1
	fmt.Scan(&T)
	n, _ := fmt.Scan(&l, &w, &h)

	for count <= T && n == 3 {
		if l <= 20 && w <= 20 && h <= 20 {
			fmt.Printf("Case %v: good\n", count)
		} else {
			fmt.Printf("Case %v: bad\n", count)
		}
		count++
		if count <= T {
			n, _ = fmt.Scan(&l, &w, &h)
		}
	}

}
